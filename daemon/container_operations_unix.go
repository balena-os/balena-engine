// +build linux freebsd

package daemon // import "github.com/docker/docker/daemon"

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/docker/docker/container"
	"github.com/docker/docker/daemon/links"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/docker/pkg/stringid"
	"github.com/docker/docker/runconfig"
	"github.com/docker/libnetwork"
	"github.com/opencontainers/selinux/go-selinux/label"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func (daemon *Daemon) setupLinkedContainers(container *container.Container) ([]string, error) {
	var env []string
	children := daemon.children(container)

	bridgeSettings := container.NetworkSettings.Networks[runconfig.DefaultDaemonNetworkMode().NetworkName()]
	if bridgeSettings == nil || bridgeSettings.EndpointSettings == nil {
		return nil, nil
	}

	for linkAlias, child := range children {
		if !child.IsRunning() {
			return nil, fmt.Errorf("Cannot link to a non running container: %s AS %s", child.Name, linkAlias)
		}

		childBridgeSettings := child.NetworkSettings.Networks[runconfig.DefaultDaemonNetworkMode().NetworkName()]
		if childBridgeSettings == nil || childBridgeSettings.EndpointSettings == nil {
			return nil, fmt.Errorf("container %s not attached to default bridge network", child.ID)
		}

		link := links.NewLink(
			bridgeSettings.IPAddress,
			childBridgeSettings.IPAddress,
			linkAlias,
			child.Config.Env,
			child.Config.ExposedPorts,
		)

		env = append(env, link.ToEnv()...)
	}

	return env, nil
}

func (daemon *Daemon) getIpcContainer(id string) (*container.Container, error) {
	errMsg := "can't join IPC of container " + id
	// Check the container exists
	container, err := daemon.GetContainer(id)
	if err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	// Check the container is running and not restarting
	if err := daemon.checkContainer(container, containerIsRunning, containerIsNotRestarting); err != nil {
		return nil, errors.Wrap(err, errMsg)
	}
	// Check the container ipc is shareable
	if st, err := os.Stat(container.ShmPath); err != nil || !st.IsDir() {
		if err == nil || os.IsNotExist(err) {
			return nil, errors.New(errMsg + ": non-shareable IPC (hint: use IpcMode:shareable for the donor container)")
		}
		// stat() failed?
		return nil, errors.Wrap(err, errMsg+": unexpected error from stat "+container.ShmPath)
	}

	return container, nil
}

func (daemon *Daemon) getPidContainer(container *container.Container) (*container.Container, error) {
	containerID := container.HostConfig.PidMode.Container()
	container, err := daemon.GetContainer(containerID)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot join PID of a non running container: %s", containerID)
	}
	return container, daemon.checkContainer(container, containerIsRunning, containerIsNotRestarting)
}

func containerIsRunning(c *container.Container) error {
	if !c.IsRunning() {
		return errdefs.Conflict(errors.Errorf("container %s is not running", c.ID))
	}
	return nil
}

func containerIsNotRestarting(c *container.Container) error {
	if c.IsRestarting() {
		return errContainerIsRestarting(c.ID)
	}
	return nil
}

func (daemon *Daemon) setupIpcDirs(c *container.Container) error {
	ipcMode := c.HostConfig.IpcMode

	switch {
	case ipcMode.IsContainer():
		ic, err := daemon.getIpcContainer(ipcMode.Container())
		if err != nil {
			return err
		}
		c.ShmPath = ic.ShmPath

	case ipcMode.IsHost():
		if _, err := os.Stat("/dev/shm"); err != nil {
			return fmt.Errorf("/dev/shm is not mounted, but must be for --ipc=host")
		}
		c.ShmPath = "/dev/shm"

	case ipcMode.IsPrivate(), ipcMode.IsNone():
		// c.ShmPath will/should not be used, so make it empty.
		// Container's /dev/shm mount comes from OCI spec.
		c.ShmPath = ""

	case ipcMode.IsEmpty():
		// A container was created by an older version of the daemon.
		// The default behavior used to be what is now called "shareable".
		fallthrough

	case ipcMode.IsShareable():
		rootIDs := daemon.idMapping.RootPair()
		if !c.HasMountFor("/dev/shm") {
			shmPath, err := c.ShmResourcePath()
			if err != nil {
				return err
			}

			if err := idtools.MkdirAllAndChown(shmPath, 0700, rootIDs); err != nil {
				return err
			}

			shmproperty := "mode=1777,size=" + strconv.FormatInt(c.HostConfig.ShmSize, 10)
			if err := unix.Mount("shm", shmPath, "tmpfs", uintptr(unix.MS_NOEXEC|unix.MS_NOSUID|unix.MS_NODEV), label.FormatMountLabel(shmproperty, c.GetMountLabel())); err != nil {
				return fmt.Errorf("mounting shm tmpfs: %s", err)
			}
			if err := os.Chown(shmPath, rootIDs.UID, rootIDs.GID); err != nil {
				return err
			}
			c.ShmPath = shmPath
		}

	default:
		return fmt.Errorf("invalid IPC mode: %v", ipcMode)
	}

	return nil
}

func (daemon *Daemon) setupSecretDir(c *container.Container) (setupErr error) {
	if len(c.SecretReferences) == 0 && len(c.ConfigReferences) == 0 {
		return nil
	}

	return fmt.Errorf("secret store is not initialized")
}

func (daemon *Daemon) cleanupSecretDir(c *container.Container) {
	// noop
}

func killProcessDirectly(cntr *container.Container) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Block until the container to stops or timeout.
	status := <-cntr.Wait(ctx, container.WaitConditionNotRunning)
	if status.Err() != nil {
		// Ensure that we don't kill ourselves
		if pid := cntr.GetPID(); pid != 0 {
			logrus.Infof("Container %s failed to exit within 10 seconds of kill - trying direct SIGKILL", stringid.TruncateID(cntr.ID))
			if err := unix.Kill(pid, 9); err != nil {
				if err != unix.ESRCH {
					return err
				}
				e := errNoSuchProcess{pid, 9}
				logrus.Debug(e)
				return e
			}
		}
	}
	return nil
}

func isLinkable(child *container.Container) bool {
	// A container is linkable only if it belongs to the default network
	_, ok := child.NetworkSettings.Networks[runconfig.DefaultDaemonNetworkMode().NetworkName()]
	return ok
}

func enableIPOnPredefinedNetwork() bool {
	return false
}

// serviceDiscoveryOnDefaultNetwork indicates if service discovery is supported on the default network
func serviceDiscoveryOnDefaultNetwork() bool {
	return false
}

func (daemon *Daemon) setupPathsAndSandboxOptions(container *container.Container, sboxOptions *[]libnetwork.SandboxOption) error {
	var err error

	if container.HostConfig.NetworkMode.IsHost() {
		// Point to the host files, so that will be copied into the container running in host mode
		*sboxOptions = append(*sboxOptions, libnetwork.OptionOriginHostsPath("/etc/hosts"))
	}

	// Copy the host's resolv.conf for the container (/etc/resolv.conf or /run/systemd/resolve/resolv.conf)
	*sboxOptions = append(*sboxOptions, libnetwork.OptionOriginResolvConfPath(daemon.configStore.GetResolvConf()))

	container.HostsPath, err = container.GetRootResourcePath("hosts")
	if err != nil {
		return err
	}
	*sboxOptions = append(*sboxOptions, libnetwork.OptionHostsPath(container.HostsPath))

	container.ResolvConfPath, err = container.GetRootResourcePath("resolv.conf")
	if err != nil {
		return err
	}
	*sboxOptions = append(*sboxOptions, libnetwork.OptionResolvConfPath(container.ResolvConfPath))
	return nil
}

func (daemon *Daemon) initializeNetworkingPaths(container *container.Container, nc *container.Container) error {
	container.HostnamePath = nc.HostnamePath
	container.HostsPath = nc.HostsPath
	container.ResolvConfPath = nc.ResolvConfPath
	return nil
}

func (daemon *Daemon) setupContainerMountsRoot(c *container.Container) error {
	// get the root mount path so we can make it unbindable
	p, err := c.MountsResourcePath("")
	if err != nil {
		return err
	}
	return idtools.MkdirAllAndChown(p, 0700, daemon.idMapping.RootPair())
}

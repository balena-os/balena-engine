package daemon // import "github.com/docker/docker/daemon"

import (
	"archive/tar"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/balena-os/librsync-go"
	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	networktypes "github.com/docker/docker/api/types/network"
	"github.com/docker/docker/container"
	"github.com/docker/docker/errdefs"
	"github.com/docker/docker/image"
	"github.com/docker/docker/pkg/idtools"
	"github.com/docker/docker/pkg/ioutils"
	"github.com/docker/docker/pkg/progress"
	"github.com/docker/docker/pkg/streamformatter"
	"github.com/docker/docker/pkg/stringid"
	"github.com/docker/docker/pkg/system"
	"github.com/docker/docker/runconfig"
	units "github.com/docker/go-units"
	"github.com/opencontainers/selinux/go-selinux/label"
	"github.com/sirupsen/logrus"
)

// CreateManagedContainer creates a container that is managed by a Service
func (daemon *Daemon) CreateManagedContainer(params types.ContainerCreateConfig) (containertypes.ContainerCreateCreatedBody, error) {
	return daemon.containerCreate(params, true)
}

// ContainerCreate creates a regular container
func (daemon *Daemon) ContainerCreate(params types.ContainerCreateConfig) (containertypes.ContainerCreateCreatedBody, error) {
	return daemon.containerCreate(params, false)
}

func (daemon *Daemon) containerCreate(params types.ContainerCreateConfig, managed bool) (containertypes.ContainerCreateCreatedBody, error) {
	start := time.Now()
	if params.Config == nil {
		return containertypes.ContainerCreateCreatedBody{}, errdefs.InvalidParameter(errors.New("Config cannot be empty in order to create a container"))
	}

	os := runtime.GOOS
	if params.Config.Image != "" {
		img, err := daemon.imageService.GetImage(params.Config.Image)
		if err == nil {
			os = img.OS
		}
	} else {
		// This mean scratch. On Windows, we can safely assume that this is a linux
		// container. On other platforms, it's the host OS (which it already is)
		if runtime.GOOS == "windows" && system.LCOWSupported() {
			os = "linux"
		}
	}

	warnings, err := daemon.verifyContainerSettings(os, params.HostConfig, params.Config, false)
	if err != nil {
		return containertypes.ContainerCreateCreatedBody{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	err = verifyNetworkingConfig(params.NetworkingConfig)
	if err != nil {
		return containertypes.ContainerCreateCreatedBody{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	if params.HostConfig == nil {
		params.HostConfig = &containertypes.HostConfig{}
	}
	err = daemon.adaptContainerSettings(params.HostConfig, params.AdjustCPUShares)
	if err != nil {
		return containertypes.ContainerCreateCreatedBody{Warnings: warnings}, errdefs.InvalidParameter(err)
	}

	container, err := daemon.create(params, managed)
	if err != nil {
		return containertypes.ContainerCreateCreatedBody{Warnings: warnings}, err
	}
	containerActions.WithValues("create").UpdateSince(start)

	return containertypes.ContainerCreateCreatedBody{ID: container.ID, Warnings: warnings}, nil
}

// Create creates a new container from the given configuration with a given name.
func (daemon *Daemon) create(params types.ContainerCreateConfig, managed bool) (retC *container.Container, retErr error) {
	var (
		container *container.Container
		img       *image.Image
		imgID     image.ID
		err       error
	)

	os := runtime.GOOS
	if params.Config.Image != "" {
		img, err = daemon.imageService.GetImage(params.Config.Image)
		if err != nil {
			return nil, err
		}
		if img.OS != "" {
			os = img.OS
		} else {
			// default to the host OS except on Windows with LCOW
			if runtime.GOOS == "windows" && system.LCOWSupported() {
				os = "linux"
			}
		}
		imgID = img.ID()

		if runtime.GOOS == "windows" && img.OS == "linux" && !system.LCOWSupported() {
			return nil, errors.New("operating system on which parent image was created is not Windows")
		}
	} else {
		if runtime.GOOS == "windows" {
			os = "linux" // 'scratch' case.
		}
	}

	if err := daemon.mergeAndVerifyConfig(params.Config, img); err != nil {
		return nil, errdefs.InvalidParameter(err)
	}

	if err := daemon.mergeAndVerifyLogConfig(&params.HostConfig.LogConfig); err != nil {
		return nil, errdefs.InvalidParameter(err)
	}

	if container, err = daemon.newContainer(params.Name, os, params.Config, params.HostConfig, imgID, managed); err != nil {
		return nil, err
	}
	defer func() {
		if retErr != nil {
			if err := daemon.cleanupContainer(container, true, true); err != nil {
				logrus.Errorf("failed to cleanup container on create error: %v", err)
			}
		}
	}()

	if err := daemon.setSecurityOptions(container, params.HostConfig); err != nil {
		return nil, err
	}

	container.HostConfig.StorageOpt = params.HostConfig.StorageOpt

	// Fixes: https://github.com/moby/moby/issues/34074 and
	// https://github.com/docker/for-win/issues/999.
	// Merge the daemon's storage options if they aren't already present. We only
	// do this on Windows as there's no effective sandbox size limit other than
	// physical on Linux.
	if runtime.GOOS == "windows" {
		if container.HostConfig.StorageOpt == nil {
			container.HostConfig.StorageOpt = make(map[string]string)
		}
		for _, v := range daemon.configStore.GraphOptions {
			opt := strings.SplitN(v, "=", 2)
			if _, ok := container.HostConfig.StorageOpt[opt[0]]; !ok {
				container.HostConfig.StorageOpt[opt[0]] = opt[1]
			}
		}
	}

	initFunc := setupInitLayer(daemon.idMapping)

	// containers that are meant to be booted from do not need the initLayer
	if params.HostConfig.Runtim == "bare" {
		initFunc = nil
	}

	// Set RWLayer for container after mount labels have been set
	rwLayer, err := daemon.imageService.CreateLayer(container, initFunc)
	if err != nil {
		return nil, errdefs.System(err)
	}
	container.RWLayer = rwLayer

	rootIDs := daemon.idMapping.RootPair()

	if err := idtools.MkdirAndChown(container.Root, 0700, rootIDs); err != nil {
		return nil, err
	}
	if err := idtools.MkdirAndChown(container.CheckpointDir(), 0700, rootIDs); err != nil {
		return nil, err
	}

	if err := daemon.setHostConfig(container, params.HostConfig); err != nil {
		return nil, err
	}

	if err := daemon.createContainerOSSpecificSettings(container, params.Config, params.HostConfig); err != nil {
		return nil, err
	}

	var endpointsConfigs map[string]*networktypes.EndpointSettings
	if params.NetworkingConfig != nil {
		endpointsConfigs = params.NetworkingConfig.EndpointsConfig
	}
	// Make sure NetworkMode has an acceptable value. We do this to ensure
	// backwards API compatibility.
	runconfig.SetDefaultNetModeIfBlank(container.HostConfig)

	daemon.updateContainerNetworkSettings(container, endpointsConfigs)
	if err := daemon.Register(container); err != nil {
		return nil, err
	}
	stateCtr.set(container.ID, "stopped")
	daemon.LogContainerEvent(container, "create")
	return container, nil
}

func toHostConfigSelinuxLabels(labels []string) []string {
	for i, l := range labels {
		labels[i] = "label=" + l
	}
	return labels
}

func (daemon *Daemon) generateSecurityOpt(hostConfig *containertypes.HostConfig) ([]string, error) {
	for _, opt := range hostConfig.SecurityOpt {
		con := strings.Split(opt, "=")
		if con[0] == "label" {
			// Caller overrode SecurityOpts
			return nil, nil
		}
	}
	ipcMode := hostConfig.IpcMode
	pidMode := hostConfig.PidMode
	privileged := hostConfig.Privileged
	if ipcMode.IsHost() || pidMode.IsHost() || privileged {
		return toHostConfigSelinuxLabels(label.DisableSecOpt()), nil
	}

	var ipcLabel []string
	var pidLabel []string
	ipcContainer := ipcMode.Container()
	pidContainer := pidMode.Container()
	if ipcContainer != "" {
		c, err := daemon.GetContainer(ipcContainer)
		if err != nil {
			return nil, err
		}
		ipcLabel = label.DupSecOpt(c.ProcessLabel)
		if pidContainer == "" {
			return toHostConfigSelinuxLabels(ipcLabel), err
		}
	}
	if pidContainer != "" {
		c, err := daemon.GetContainer(pidContainer)
		if err != nil {
			return nil, err
		}

		pidLabel = label.DupSecOpt(c.ProcessLabel)
		if ipcContainer == "" {
			return toHostConfigSelinuxLabels(pidLabel), err
		}
	}

	if pidLabel != nil && ipcLabel != nil {
		for i := 0; i < len(pidLabel); i++ {
			if pidLabel[i] != ipcLabel[i] {
				return nil, fmt.Errorf("--ipc and --pid containers SELinux labels aren't the same")
			}
		}
		return toHostConfigSelinuxLabels(pidLabel), nil
	}
	return nil, nil
}

func (daemon *Daemon) mergeAndVerifyConfig(config *containertypes.Config, img *image.Image) error {
	if img != nil && img.Config != nil {
		if err := merge(config, img.Config); err != nil {
			return err
		}
	}
	// Reset the Entrypoint if it is [""]
	if len(config.Entrypoint) == 1 && config.Entrypoint[0] == "" {
		config.Entrypoint = nil
	}
	if len(config.Entrypoint) == 0 && len(config.Cmd) == 0 {
		return fmt.Errorf("No command specified")
	}
	return nil
}

// Checks if the client set configurations for more than one network while creating a container
// Also checks if the IPAMConfig is valid
func verifyNetworkingConfig(nwConfig *networktypes.NetworkingConfig) error {
	if nwConfig == nil || len(nwConfig.EndpointsConfig) == 0 {
		return nil
	}
	if len(nwConfig.EndpointsConfig) == 1 {
		for k, v := range nwConfig.EndpointsConfig {
			if v == nil {
				return errdefs.InvalidParameter(errors.Errorf("no EndpointSettings for %s", k))
			}
			if v.IPAMConfig != nil {
				if v.IPAMConfig.IPv4Address != "" && net.ParseIP(v.IPAMConfig.IPv4Address).To4() == nil {
					return errors.Errorf("invalid IPv4 address: %s", v.IPAMConfig.IPv4Address)
				}
				if v.IPAMConfig.IPv6Address != "" {
					n := net.ParseIP(v.IPAMConfig.IPv6Address)
					// if the address is an invalid network address (ParseIP == nil) or if it is
					// an IPv4 address (To4() != nil), then it is an invalid IPv6 address
					if n == nil || n.To4() != nil {
						return errors.Errorf("invalid IPv6 address: %s", v.IPAMConfig.IPv6Address)
					}
				}
			}
		}
		return nil
	}
	l := make([]string, 0, len(nwConfig.EndpointsConfig))
	for k := range nwConfig.EndpointsConfig {
		l = append(l, k)
	}
	return errors.Errorf("Container cannot be connected to network endpoints: %s", strings.Join(l, ", "))
}

// DeltaCreate creates a delta of the specified src and dest images
// This is called directly from the Engine API
func (daemon *Daemon) DeltaCreate(deltaSrc, deltaDest string, options types.ImageDeltaOptions, outStream io.Writer) error {
	progressOutput := streamformatter.NewJSONProgressOutput(outStream, false)

	srcImg, err := daemon.GetImage(deltaSrc)
	if err != nil {
		return errors.Wrapf(err, "no such image: %s", deltaSrc)
	}

	dstImg, err := daemon.GetImage(deltaDest)
	if err != nil {
		return errors.Wrapf(err, "no such image: %s", deltaDest)
	}

	is := daemon.stores[dstImg.OperatingSystem()].imageStore
	ls := daemon.stores[dstImg.OperatingSystem()].layerStore

	srcData, err := is.GetTarSeekStream(srcImg.ID())
	if err != nil {
		return err
	}
	defer srcData.Close()

	srcDataLen, err := ioutils.SeekerSize(srcData)
	if err != nil {
		return err
	}

	progressReader := progress.NewProgressReader(srcData, progressOutput, srcDataLen, deltaSrc, "Fingerprinting")
	defer progressReader.Close()

	srcSig, err := librsync.Signature(bufio.NewReaderSize(progressReader, 65536), ioutil.Discard, 512, 32, librsync.BLAKE2_SIG_MAGIC)
	if err != nil {
		return err
	}

	progress.Update(progressOutput, deltaSrc, "Fingerprint complete")

	deltaRootFS := image.NewRootFS()

	for _, diffID := range dstImg.RootFS.DiffIDs {
		progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Waiting")
	}

	statTotalSize := int64(0)
	statDeltaSize := int64(0)

	for i, diffID := range dstImg.RootFS.DiffIDs {
		var (
			layerData io.Reader
			platform  layer.OS
		)

		commonLayer := false
		dstRootFS := *dstImg.RootFS
		dstRootFS.DiffIDs = dstRootFS.DiffIDs[:i+1]

		if i < len(srcImg.RootFS.DiffIDs) {
			srcRootFS := *srcImg.RootFS
			srcRootFS.DiffIDs = srcRootFS.DiffIDs[:i+1]

			if srcRootFS.ChainID() == dstRootFS.ChainID() {
				commonLayer = true
			}
		}

		// We're only interested in layers that are different. Push empty
		// layers for common layers
		if commonLayer {
			layerData, _ = layer.EmptyLayer.TarStream()
			platform = layer.EmptyLayer.OS()
		} else {

			l, err := ls.Get(dstRootFS.ChainID())
			if err != nil {
				return err
			}
			defer layer.ReleaseAndLog(ls, l)

			platform = l.OS()

			input, err := l.TarStream()
			if err != nil {
				return err
			}
			defer input.Close()

			inputSize, err := l.DiffSize()
			if err != nil {
				return err
			}

			statTotalSize += inputSize

			progressReader := progress.NewProgressReader(input, progressOutput, inputSize, stringid.TruncateID(diffID.String()), "Computing delta")
			defer progressReader.Close()

			pR, pW := io.Pipe()

			layerData = pR

			tmpDelta, err := ioutil.TempFile("", "docker-delta-")
			if err != nil {
				return err
			}
			defer os.Remove(tmpDelta.Name())

			go func() {
				w := bufio.NewWriter(tmpDelta)
				err := librsync.Delta(srcSig, bufio.NewReader(progressReader), w)
				if err != nil {
					pW.CloseWithError(err)
					return
				}
				w.Flush()

				info, err := tmpDelta.Stat()
				if err != nil {
					pW.CloseWithError(err)
					return
				}

				tw := tar.NewWriter(pW)

				hdr := &tar.Header{
					Name: "delta",
					Mode: 0600,
					Size: info.Size(),
				}

				if err := tw.WriteHeader(hdr); err != nil {
					pW.CloseWithError(err)
					return
				}

				if _, err := tmpDelta.Seek(0, io.SeekStart); err != nil {
					pW.CloseWithError(err)
					return
				}

				if _, err := io.Copy(tw, tmpDelta); err != nil {
					pW.CloseWithError(err)
					return
				}

				if err := tw.Close(); err != nil {
					pW.CloseWithError(err)
					return
				}

				pW.Close()
			}()
		}

		newLayer, err := ls.Register(layerData, deltaRootFS.ChainID(), platform)
		if err != nil {
			return err
		}
		defer layer.ReleaseAndLog(ls, newLayer)

		if commonLayer {
			progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Skipping common layer")
		} else {
			deltaSize, err := newLayer.DiffSize()
			if err != nil {
				return err
			}
			statDeltaSize += deltaSize
			progress.Update(progressOutput, stringid.TruncateID(diffID.String()), "Delta complete")
		}

		deltaRootFS.Append(newLayer.DiffID())
	}

	config := image.Image{
		RootFS: deltaRootFS,
		V1Image: image.V1Image{
			Created: time.Now().UTC(),
			Config: &containertypes.Config{
				Labels: map[string]string{
					"io.resin.delta.base":   srcImg.ID().String(),
					"io.resin.delta.config": string(dstImg.RawJSON()),
				},
			},
		},
	}

	rawConfig, err := json.Marshal(config)
	if err != nil {
		return err
	}

	id, err := is.Create(rawConfig)
	if err != nil {
		return err
	}

	humanTotal := units.HumanSize(float64(statTotalSize))
	humanDelta := units.HumanSize(float64(statDeltaSize))
	deltaRatio := float64(statTotalSize) / float64(statDeltaSize)
	if statTotalSize == 0 {
		deltaRatio = 1
	}

	outStream.Write(streamformatter.FormatStatus("", "Normal size: %s, Delta size: %s, %.2fx improvement", humanTotal, humanDelta, deltaRatio))
	outStream.Write(streamformatter.FormatStatus("", "Created delta: %s", id.String()))

	if options.Tag == "" {
		return nil
	}

	ref, err := reference.ParseNormalizedNamed(options.Tag)
	if err != nil {
		return err
	}

	if _, isCanonical := ref.(reference.Canonical); isCanonical {
		return errors.New("build tag cannot contain a digest")
	}

	ref = reference.TagNameOnly(ref)

	if err := daemon.TagImageWithReference(id, runtime.GOOS, ref); err != nil {
		return err
	}
	logrus.Debugf("Tagged delta %s with %s", id.String(), reference.FamiliarString(ref))
	outStream.Write(streamformatter.FormatStatus("", "Successfully tagged %s\n", reference.FamiliarString(ref)))

	return nil
}

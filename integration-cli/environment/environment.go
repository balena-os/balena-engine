package environment

import (
	"os"

	"os/exec"

	"github.com/docker/docker/internal/test/environment"
)

var (
	// DefaultClientBinary is the name of the docker binary
	DefaultClientBinary = os.Getenv("TEST_CLIENT_BINARY")
)

func init() {
	if DefaultClientBinary == "" {
		DefaultClientBinary = "docker"
	}
}

// Execution contains information about the current test execution and daemon
// under test
type Execution struct {
	environment.Execution
	dockerBinary string
}

// DockerBinary returns the docker binary for this testing environment
func (e *Execution) DockerBinary() string {
	return e.dockerBinary
}

// New returns details about the testing environment
func New() (*Execution, error) {
	env, err := environment.New()
	if err != nil {
		return nil, err
	}
	daemonPlatform := info.OSType
	if daemonPlatform != "linux" && daemonPlatform != "windows" {
		return nil, fmt.Errorf("Cannot run tests against platform: %s", daemonPlatform)
	}
	baseImage := "scratch"
	volumesConfigPath := filepath.Join(info.DockerRootDir, "volumes")
	containerStoragePath := filepath.Join(info.DockerRootDir, "containers")
	// Make sure in context of daemon, not the local platform. Note we can't
	// use filepath.FromSlash or ToSlash here as they are a no-op on Unix.
	if daemonPlatform == "windows" {
		volumesConfigPath = strings.Replace(volumesConfigPath, `/`, `\`, -1)
		containerStoragePath = strings.Replace(containerStoragePath, `/`, `\`, -1)

		baseImage = "microsoft/windowsservercore"
		if len(os.Getenv("WINDOWS_BASE_IMAGE")) > 0 {
			baseImage = os.Getenv("WINDOWS_BASE_IMAGE")
			fmt.Println("INFO: Windows Base image is ", baseImage)
		}
	} else {
		volumesConfigPath = strings.Replace(volumesConfigPath, `\`, `/`, -1)
		containerStoragePath = strings.Replace(containerStoragePath, `\`, `/`, -1)
	}

	var daemonPid int
	dest := os.Getenv("DEST")
	b, err := ioutil.ReadFile(filepath.Join(dest, "balena.pid"))
	if err == nil {
		if p, err := strconv.ParseInt(string(b), 10, 32); err == nil {
			daemonPid = int(p)
		}
	}
	dockerBinary, err := exec.LookPath(DefaultClientBinary)
	if err != nil {
		return nil, err
	}

	return &Execution{
		Execution:    *env,
		dockerBinary: dockerBinary,
	}, nil
}

// DockerBasePath is the base path of the docker folder (by default it is -/var/run/docker)
// TODO: remove
// Deprecated: use Execution.DaemonInfo.DockerRootDir
func (e *Execution) DockerBasePath() string {
	return e.DaemonInfo.DockerRootDir
}

// ExperimentalDaemon tell whether the main daemon has
// experimental features enabled or not
// Deprecated: use DaemonInfo.ExperimentalBuild
func (e *Execution) ExperimentalDaemon() bool {
	return e.DaemonInfo.ExperimentalBuild
}

// DaemonPlatform is held globally so that tests can make intelligent
// decisions on how to configure themselves according to the platform
// of the daemon. This is initialized in docker_utils by sending
// a version call to the daemon and examining the response header.
// Deprecated: use Execution.OSType
func (e *Execution) DaemonPlatform() string {
	return e.OSType
}

// MinimalBaseImage is the image used for minimal builds (it depends on the platform)
// Deprecated: use Execution.PlatformDefaults.BaseImage
func (e *Execution) MinimalBaseImage() string {
	return e.PlatformDefaults.BaseImage
}

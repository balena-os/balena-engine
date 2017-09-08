package containerdShim

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/containerd/containerd/osutils"
	"github.com/containerd/containerd/specs"
	"github.com/containerd/console"
	runc "github.com/containerd/go-runc"
)

var errRuntime = errors.New("shim: runtime execution error")

type checkpoint struct {
	// Timestamp is the time that checkpoint happened
	Created time.Time `json:"created"`
	// Name is the name of the checkpoint
	Name string `json:"name"`
	// TCP checkpoints open tcp connections
	TCP bool `json:"tcp"`
	// UnixSockets persists unix sockets in the checkpoint
	UnixSockets bool `json:"unixSockets"`
	// Shell persists tty sessions in the checkpoint
	Shell bool `json:"shell"`
	// Exit exits the container after the checkpoint is finished
	Exit bool `json:"exit"`
	// EmptyNS tells CRIU not to restore a particular namespace
	EmptyNS []string `json:"emptyNS,omitempty"`
}

type processState struct {
	specs.ProcessSpec
	Exec           bool     `json:"exec"`
	Stdin          string   `json:"containerdStdin"`
	Stdout         string   `json:"containerdStdout"`
	Stderr         string   `json:"containerdStderr"`
	RuntimeArgs    []string `json:"runtimeArgs"`
	NoPivotRoot    bool     `json:"noPivotRoot"`
	CheckpointPath string   `json:"checkpoint"`
	RootUID        int      `json:"rootUID"`
	RootGID        int      `json:"rootGID"`
}

type process struct {
	sync.WaitGroup
	id             string
	bundle         string
	stdio          *stdio
	containerPid   int
	checkpoint     *checkpoint
	checkpointPath string
	shimIO         *IO
	stdinCloser    io.Closer
	state          *processState
	runtime        string
	ioCleanupFn    func()

	socket       *runc.Socket
	console      console.Console
	consoleErrCh chan error
}

func newProcess(id, bundle, runtimeName string) (*process, error) {
	p := &process{
		id:           id,
		bundle:       bundle,
		runtime:      runtimeName,
		consoleErrCh: make(chan error, 1),
	}
	s, err := loadProcess()
	if err != nil {
		return nil, err
	}
	p.state = s
	if s.CheckpointPath != "" {
		cpt, err := loadCheckpoint(s.CheckpointPath)
		if err != nil {
			return nil, err
		}
		p.checkpoint = cpt
		p.checkpointPath = s.CheckpointPath
	}
	return p, nil
}

func loadProcess() (*processState, error) {
	f, err := os.Open("process.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var s processState
	if err := json.NewDecoder(f).Decode(&s); err != nil {
		return nil, err
	}
	return &s, nil
}

func loadCheckpoint(checkpointPath string) (*checkpoint, error) {
	f, err := os.Open(filepath.Join(checkpointPath, "config.json"))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cpt checkpoint
	if err := json.NewDecoder(f).Decode(&cpt); err != nil {
		return nil, err
	}
	return &cpt, nil
}

func (p *process) create() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	logPath := filepath.Join(cwd, "log.json")
	args := append([]string{
		"--log", logPath,
		"--log-format", "json",
	}, p.state.RuntimeArgs...)
	if p.state.Exec {
		args = append(args, "exec",
			"-d",
			"--process", filepath.Join(cwd, "process.json"),
		)
		if p.socket != nil {
			args = append(args, "--console-socket", p.socket.Path())
		}

	} else if p.checkpoint != nil {
		args = append(args, "restore",
			"-d",
			"--image-path", p.checkpointPath,
			"--work-path", filepath.Join(p.checkpointPath, "criu.work", "restore-"+time.Now().Format(time.RFC3339Nano)),
		)
		add := func(flags ...string) {
			args = append(args, flags...)
		}
		if p.checkpoint.Shell {
			add("--shell-job")
		}
		if p.checkpoint.TCP {
			add("--tcp-established")
		}
		if p.checkpoint.UnixSockets {
			add("--ext-unix-sk")
		}
		if p.state.NoPivotRoot {
			add("--no-pivot")
		}
		for _, ns := range p.checkpoint.EmptyNS {
			add("--empty-ns", ns)
		}

	} else {
		args = append(args, "create",
			"--bundle", p.bundle,
		)
		if p.socket != nil {
			args = append(args, "--console-socket", p.socket.Path())
		}
		if p.state.NoPivotRoot {
			args = append(args, "--no-pivot")
		}
	}
	args = append(args,
		"--pid-file", filepath.Join(cwd, "pid"),
		p.id,
	)
	cmd := exec.Command(p.runtime, args...)
	cmd.Dir = p.bundle
	cmd.Stdin = p.stdio.stdin
	cmd.Stdout = p.stdio.stdout
	cmd.Stderr = p.stdio.stderr
	// Call out to SetPDeathSig to set SysProcAttr as elements are platform specific
	cmd.SysProcAttr = osutils.SetPDeathSig()

	if err := cmd.Start(); err != nil {
		if exErr, ok := err.(*exec.Error); ok {
			if exErr.Err == exec.ErrNotFound || exErr.Err == os.ErrNotExist {
				return fmt.Errorf("%s not installed on system", p.runtime)
			}
		}
		return err
	}
	if runtime.GOOS != "solaris" {
		// Since current logic dictates that we need a pid at the end of p.create
		// we need to call runtime start as well on Solaris hence we need the
		// pipes to stay open.
		p.stdio.stdout.Close()
		p.stdio.stderr.Close()
	}
	if err := cmd.Wait(); err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return errRuntime
		}
		return err
	}
	data, err := ioutil.ReadFile("pid")
	if err != nil {
		return err
	}
	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	p.containerPid = pid
	return nil
}

func (p *process) pid() int {
	return p.containerPid
}

func (p *process) delete() error {
	if !p.state.Exec {
		cmd := exec.Command(p.runtime, append(p.state.RuntimeArgs, "delete", p.id)...)
		cmd.SysProcAttr = osutils.SetPDeathSig()
		out, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("%s: %v", out, err)
		}
	}
	return nil
}

// IO holds all 3 standard io Reader/Writer (stdin,stdout,stderr)
type IO struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser
}

func (p *process) initializeIO(rootuid, rootgid int) (i *IO, err error) {
	var fds []uintptr
	i = &IO{}
	// cleanup in case of an error
	defer func() {
		if err != nil {
			for _, fd := range fds {
				syscall.Close(int(fd))
			}
		}
	}()
	// STDIN
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stdin, i.Stdin = r, w
	// STDOUT
	if r, w, err = os.Pipe(); err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stdout, i.Stdout = w, r
	// STDERR
	if r, w, err = os.Pipe(); err != nil {
		return nil, err
	}
	fds = append(fds, r.Fd(), w.Fd())
	p.stdio.stderr, i.Stderr = w, r
	// change ownership of the pipes in case we are in a user namespace
	for _, fd := range fds {
		if err := syscall.Fchown(int(fd), rootuid, rootgid); err != nil {
			return nil, err
		}
	}
	return i, nil
}
func (p *process) Close() error {
	err := p.stdio.Close()
	if p.socket != nil {
		p.socket.Close()
	}
	return err
}

type stdio struct {
	stdin  *os.File
	stdout *os.File
	stderr *os.File
}

func (s *stdio) Close() error {
	err := s.stdin.Close()
	if oerr := s.stdout.Close(); err == nil {
		err = oerr
	}
	if oerr := s.stderr.Close(); err == nil {
		err = oerr
	}
	return err
}

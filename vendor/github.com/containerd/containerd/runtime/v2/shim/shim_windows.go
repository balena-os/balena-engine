// +build windows

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package shim

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"sync"

	winio "github.com/Microsoft/go-winio"
	"github.com/containerd/containerd/events"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/ttrpc"
	"github.com/containerd/typeurl"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// setupSignals creates a new signal handler for all signals
func setupSignals() (chan os.Signal, error) {
	signals := make(chan os.Signal, 32)
	return signals, nil
}

func newServer() (*ttrpc.Server, error) {
	return ttrpc.NewServer()
}

func subreaper() error {
	return nil
}

func setupDumpStacks(dump chan<- os.Signal) {
	// TODO: JTERRY75: Make this based on events. signal.Notify(dump, syscall.SIGUSR1)
}

// serve serves the ttrpc API over a unix socket at the provided path
// this function does not block
func serveListener(path string) (net.Listener, error) {
	if path == "" {
		return nil, errors.New("'socket' must be npipe path")
	}
	l, err := winio.ListenPipe(path, nil)
	if err != nil {
		return nil, err
	}
	logrus.WithField("socket", path).Debug("serving api on npipe socket")
	return l, nil
}

func handleSignals(logger *logrus.Entry, signals chan os.Signal) error {
	logger.Info("starting signal loop")
	for {
		select {
		case s := <-signals:
			switch s {
			case os.Interrupt:
				break
			}
		}
	}
}

type deferredShimWriteLogger struct {
	ctx context.Context

	wg sync.WaitGroup

	c      net.Conn
	conerr error
}

func (dswl *deferredShimWriteLogger) Write(p []byte) (int, error) {
	dswl.wg.Wait()
	if dswl.c == nil {
		return 0, dswl.conerr
	}
	return dswl.c.Write(p)
}

// openLog on Windows acts as the server of the log pipe. This allows the
// containerd daemon to independently restart and reconnect to the logs.
func openLog(ctx context.Context, id string) (io.Writer, error) {
	ns, err := namespaces.NamespaceRequired(ctx)
	if err != nil {
		return nil, err
	}
	l, err := winio.ListenPipe(fmt.Sprintf("\\\\.\\pipe\\containerd-shim-%s-%s-log", ns, id), nil)
	if err != nil {
		return nil, err
	}
	dswl := &deferredShimWriteLogger{
		ctx: ctx,
	}
	// TODO: JTERRY75 - this will not work with restarts. Only the first
	// connection will work and all +1 connections will return 'use of closed
	// network connection'. Make this reconnect aware.
	dswl.wg.Add(1)
	go func() {
		c, conerr := l.Accept()
		if conerr != nil {
			l.Close()
			dswl.conerr = conerr
		}
		dswl.c = c
		dswl.wg.Done()
	}()
	return dswl, nil
}

func (l *remoteEventsPublisher) Publish(ctx context.Context, topic string, event events.Event) error {
	ns, _ := namespaces.Namespace(ctx)
	encoded, err := typeurl.MarshalAny(event)
	if err != nil {
		return err
	}
	data, err := encoded.Marshal()
	if err != nil {
		return err
	}
	cmd := exec.CommandContext(ctx, l.containerdBinaryPath, "--address", l.address, "publish", "--topic", topic, "--namespace", ns)
	cmd.Stdin = bytes.NewReader(data)
	return cmd.Run()
}

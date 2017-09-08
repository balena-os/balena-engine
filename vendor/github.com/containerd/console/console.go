package console

import (
	"errors"
	"io"
	"os"
)

var ErrNotAConsole = errors.New("provided file is not a console")

type Console interface {
	io.Reader
	io.Writer
	io.Closer

	// Resize resizes the console to the provided window size
	Resize(WinSize) error
	// ResizeFrom resizes the calling console to the size of the
	// provided console
	ResizeFrom(Console) error
	// SetRaw sets the console in raw mode
	SetRaw() error
	// DisableEcho disables echo on the console
	DisableEcho() error
	// Reset restores the console to its orignal state
	Reset() error
	// Size returns the window size of the console
	Size() (WinSize, error)
}

// WinSize specifies the window size of the console
type WinSize struct {
	// Height of the console
	Height uint16
	// Width of the console
	Width uint16
	x     uint16
	y     uint16
}

// Current returns the current processes console
func Current() Console {
	return newMaster(os.Stdin)
}

// ConsoleFromFile returns a console using the provided file
func ConsoleFromFile(f *os.File) (Console, error) {
	if err := checkConsole(f); err != nil {
		return nil, err
	}
	return newMaster(f), nil
}

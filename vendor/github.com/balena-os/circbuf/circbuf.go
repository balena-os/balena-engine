package circbuf

import (
	"fmt"
)

// Buffer implements a circular buffer. It is a fixed size,
// and new writes overwrite older data, such that for a buffer
// of size N, for any amount of writes, only the last N bytes
// are retained.
type Buffer struct {
	data        []byte
	out         []byte
	size        int64
	writeCursor int64
	written     int64
}

// NewBuffer creates a new buffer of a given size. The size
// must be greater than 0.
func NewBuffer(size int64) (*Buffer, error) {
	if size <= 0 {
		return nil, fmt.Errorf("Size must be positive")
	}

	b := &Buffer{
		size: size,
		data: make([]byte, size),
		out:  make([]byte, size),
	}
	return b, nil
}

// Write writes up to len(buf) bytes to the internal ring,
// overriding older data if necessary.
func (b *Buffer) Write(buf []byte) (int, error) {
	// Account for total bytes written
	n := len(buf)
	b.written += int64(n)

	// If the buffer is larger than ours, then we only care
	// about the last size bytes anyways
	if int64(n) > b.size {
		buf = buf[int64(n)-b.size:]
	}

	// Copy in place
	remain := b.size - b.writeCursor
	copy(b.data[b.writeCursor:], buf)
	if int64(len(buf)) > remain {
		copy(b.data, buf[remain:])
	}

	// Update location of the cursor
	b.writeCursor = ((b.writeCursor + int64(len(buf))) % b.size)
	return n, nil
}

// WriteByte writes a single byte into the buffer.
func (b *Buffer) WriteByte(c byte) error {
	b.data[b.writeCursor] = c
	b.writeCursor = ((b.writeCursor + 1) % b.size)
	b.written++
	return nil
}

// Size returns the size of the buffer
func (b *Buffer) Size() int64 {
	return b.size
}

// TotalWritten provides the total number of bytes written
func (b *Buffer) TotalWritten() int64 {
	return b.written
}

// Bytes provides a slice of the bytes written. This
// slice should not be written to. The underlying array
// may point to data that will be overwritten by a subsequent
// call to Bytes. It does no allocation.
func (b *Buffer) Bytes() []byte {
	switch {
	case b.written >= b.size && b.writeCursor == 0:
		return b.data
	case b.written > b.size:
		copy(b.out, b.data[b.writeCursor:])
		copy(b.out[b.size-b.writeCursor:], b.data[:b.writeCursor])
		return b.out
	default:
		return b.data[:b.writeCursor]
	}
}

// Get returns a single byte out of the buffer, at the given position.
func (b *Buffer) Get(i int64) (byte, error) {
	switch {
	case i >= b.written || i >= b.size:
		return 0, fmt.Errorf("Index out of bounds: %v", i)
	case b.written > b.size:
		return b.data[(b.writeCursor+i)%b.size], nil
	default:
		return b.data[i], nil
	}
}

// Reset resets the buffer so it has no content.
func (b *Buffer) Reset() {
	b.writeCursor = 0
	b.written = 0
}

// String returns the contents of the buffer as a string
func (b *Buffer) String() string {
	return string(b.Bytes())
}

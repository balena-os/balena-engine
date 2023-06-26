package ioutils

import (
	"io"
	"strconv"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

// nopCloser wraps an io.ReadSeeker, providing it with a no-op Close method.
// Like io.NopCloser, but for io.ReadSeekers.
type nopCloser struct {
	io.ReadSeeker
}

func (nopCloser) Close() error { return nil }

// Iterates, reading all data from a ConcatReadSeekCloser, using a buffer of a
// given size. Checks if no unexpected errors are generated and if all expected
// data is read.
func TestConcatReadSeekCloserRead(t *testing.T) {
	tests := []struct {
		aContents  string
		bContents  string
		bufferSize int
	}{
		// Buffers fitting the concatenated readers perfectly
		{"abcd", "1234", 4},
		{"abcdef", "123456789", 3},
		{"abc", "123", 6},

		// Buffers _not_ fitting the concatenated readers perfectly
		{"abc", "12", 2},
		{"ab", "123", 2},
		{"abcde", "123456", 3},

		// Buffers that are larger than either or both concatenated readers
		{"abc", "def", 10},
		{"abcdefgh", "1", 2},
		{"abcdefgh", "1", 3},
		{"abcdefgh", "1", 8},
		{"a", "12345678", 2},
		{"a", "12345678", 3},
		{"a", "12345678", 8},

		// Either or both readers empty
		{"", "1234", 4},
		{"abcd", "", 4},
		{"", "1234", 3},
		{"abcd", "", 3},
		{"", "", 4},
	}

	for _, tt := range tests {
		testName := tt.aContents + "/" + tt.bContents + "/" + strconv.Itoa(tt.bufferSize)
		t.Run(testName, func(t *testing.T) {
			// Creates two Readers, a and b, with the given contents, and use
			// them to create the ConcatReadSeekCloser c (the one to be tested).
			a := nopCloser{strings.NewReader(tt.aContents)}
			b := nopCloser{strings.NewReader(tt.bContents)}

			c, err := ConcatReadSeekClosers(a, b)
			assert.NilError(t, err, "error creating ConcatReadSeekCloser")

			// Read until EOF.
			buf := make([]byte, tt.bufferSize)
			readData := ""

			for {
				n, err := c.Read(buf)

				readData += string(buf[:n])
				if err == io.EOF {
					break
				}
				assert.NilError(t, err, "unexpected read error")
			}

			// Ensure we got all the data that was expected.
			expectedReadData := tt.aContents + tt.bContents
			assert.Equal(t, expectedReadData, readData, "got bad data")
		})
	}
}

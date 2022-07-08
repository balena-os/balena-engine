package ioutils

import (
	"errors"
	"io"
)

// clampSliceIndex clamps index to be between min and max (inclusive). As the
// name suggests, the function is designed to be especially useful for slice
// indices -- and in particular, to do so safely in both 32- and 64-bit
// platforms.
func clampSliceIndex(index int64, min, max int) int {
	if index < int64(min) {
		return min
	}
	if index > int64(max) {
		return max
	}
	return int(index)
}

// max64 returns the larger of two int64 values.
func max64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func SeekerSize(s io.Seeker) (int64, error) {
	cur, err := s.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	size, err := s.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	if _, err := s.Seek(cur, io.SeekStart); err != nil {
		return 0, err
	}

	return size, nil
}

type concatReadSeekCloser struct {
	a     ReadSeekCloser
	aSize int64
	b     ReadSeekCloser
	bSize int64
	off   int64
}

func (self *concatReadSeekCloser) Read(p []byte) (n int, err error) {
	// if the read starts within a
	if self.off < self.aSize {
		if _, err := self.a.Seek(self.off, io.SeekStart); err != nil {
			return 0, err
		}

		i := clampSliceIndex(self.aSize-self.off, 0, len(p))
		nA, err := io.ReadFull(self.a, p[:i])

		if err != nil {
			return 0, err
		}
		n += nA
	}

	// if the read ends within b
	if self.off+int64(len(p)) >= self.aSize {
		bOffset := max64(self.off-self.aSize, 0)

		if _, err := self.b.Seek(bOffset, io.SeekStart); err != nil {
			return 0, err
		}

		i := clampSliceIndex(self.aSize-self.off, 0, len(p))
		j := clampSliceIndex(int64(i)+self.bSize-bOffset, i, len(p))

		if i != j {
			nB, err := io.ReadFull(self.b, p[i:j])
			if err != nil {
				return 0, err
			}
			n += nB
		}
	}

	self.off += int64(n)
	if self.off == self.aSize+self.bSize {
		err = io.EOF
	}

	return
}

func (self *concatReadSeekCloser) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	default:
		return 0, errors.New("Seek: invalid whence")
	case io.SeekStart:
		break
	case io.SeekCurrent:
		offset += self.off
	case io.SeekEnd:
		offset += self.aSize + self.bSize
	}
	if offset < 0 {
		return 0, errors.New("Seek: invalid offset")
	}
	self.off = offset
	return offset, nil
}

func (self *concatReadSeekCloser) Close() error {
	err1 := self.a.Close()
	err2 := self.b.Close()

	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func ConcatReadSeekClosers(a, b ReadSeekCloser) (ReadSeekCloser, error) {
	aSize, err := SeekerSize(a)
	if err != nil {
		return nil, err
	}

	bSize, err := SeekerSize(b)
	if err != nil {
		return nil, err
	}

	return &concatReadSeekCloser{a, aSize, b, bSize, 0}, nil
}

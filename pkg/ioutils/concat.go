package ioutils

import (
	"io"
	"errors"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
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
	a ReadSeekCloser
	aSize int64
	b ReadSeekCloser
	bSize int64
	off int64
}

func (self *concatReadSeekCloser) Read(p []byte) (n int, err error) {
	// if the read starts within a
	if self.off < self.aSize {
		if _, err := self.a.Seek(self.off, io.SeekStart); err != nil {
			return 0, err
		}

		nA, err := io.ReadFull(self.a, p[:min(len(p), int(self.aSize - self.off))])
		if err != nil {
			return 0, err
		}
		n += nA
	}

	// if the read ends within b
	if self.off + int64(len(p)) >= self.aSize {
		if _, err := self.b.Seek(int64(max(0, int(self.off - self.aSize))), io.SeekStart); err != nil {
			return 0, err
		}

		nB, err := io.ReadFull(self.b, p[max(int(self.aSize - self.off), 0):])
		if err != nil {
			return 0, err
		}
		n += nB
	}

	self.off += int64(n)

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

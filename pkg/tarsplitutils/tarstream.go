package tarsplitutils

import (
	"fmt"
	"io"
	"io/ioutil"
	"sort"

	"github.com/vbatts/tar-split/tar/storage"
)

func min(x, y int64) int64 {
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

type randomAccessTarStream struct {
	entries      storage.Entries
	entryOffsets []int64
	fg           storage.FileGetter
}

func (self randomAccessTarStream) ReadAt(p []byte, off int64) (int, error) {
	// Find the first entry that we're interested in
	firstEntry := sort.Search(len(self.entryOffsets), func(i int) bool { return self.entryOffsets[i] > off }) - 1

	// The cursor will most likely be negative the first time. This signifies
	// that we need to read some data first before starting to fill the buffer
	n := self.entryOffsets[firstEntry] - off

	for _, entry := range self.entries[firstEntry:] {
		if n >= int64(len(p)) {
			break
		}

		switch entry.Type {
		case storage.SegmentType:
			payload := entry.Payload
			if n < 0 {
				payload = payload[-n:]
				n = 0
			}

			n += int64(copy(p[n:], payload))
		case storage.FileType:
			if entry.Size == 0 {
				continue
			}

			fh, err := self.fg.Get(entry.GetName())
			if err != nil {
				return 0, err
			}

			end := min(n+entry.Size, int64(len(p)))

			if n < 0 {
				if seeker, ok := fh.(io.Seeker); ok {
					if _, err := seeker.Seek(-n, io.SeekStart); err != nil {
						return 0, err
					}
				} else {
					if _, err := io.CopyN(ioutil.Discard, fh, -n); err != nil {
						return 0, err
					}
				}
				n = 0
			}

			_, err = io.ReadFull(fh, p[n:end])

			n += end - n
			if err != nil {
				return 0, err
			}

			fh.Close()
		default:
			return 0, fmt.Errorf("Unknown tar-split entry type: %v", entry.Type)
		}
	}

	return len(p), nil
}

func NewRandomAccessTarStream(fg storage.FileGetter, up storage.Unpacker) (io.ReadSeeker, error) {
	stream := randomAccessTarStream{
		fg: fg,
	}

	size := int64(0)
	for {
		entry, err := up.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		stream.entryOffsets = append(stream.entryOffsets, size)
		stream.entries = append(stream.entries, *entry)

		switch entry.Type {
		case storage.SegmentType:
			size += int64(len(entry.Payload))
		case storage.FileType:
			size += entry.Size
		}
	}

	// Push ending offset. This is because when we binary search we search for
	// the offset that is above the target and then we move one step back.
	// See implementation of ReadAt()
	stream.entryOffsets = append(stream.entryOffsets, size)

	return io.NewSectionReader(stream, 0, size), nil
}

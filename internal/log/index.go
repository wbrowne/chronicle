package log

import (
	"fmt"
	"io"
	"os"

	"github.com/tysontate/gommap"
)

var (
	offWidth uint64 = 4
	posWidth uint64 = 8
	entWidth        = offWidth + posWidth
)

type index struct {
	file *os.File
	mmap gommap.MMap
	size uint64
}

func newIndex(f *os.File, c *Config) (*index, error) {
	fmt.Println("### New index created ###")
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if err = os.Truncate(f.Name(), int64(c.Segment.MaxIndexBytes)); err != nil {
		return nil, err
	}

	idx := &index{
		file: f,
		size: uint64(fi.Size()),
	}
	if idx.mmap, err = gommap.Map(
		idx.file.Fd(),
		gommap.PROT_READ|gommap.PROT_WRITE,
		gommap.MAP_SHARED,
	); err != nil {
		return nil, err
	}

	return idx, nil
}

// out represents the relative offset to the segments offset
func (i *index) Read(in int64) (off uint32, pos uint64, err error) {
	if i.size == 0 { // empty index
		return 0, 0, io.EOF
	}

	if in == -1 {
		fmt.Printf("Current size = %d\n", i.size)
		off = uint32((i.size / entWidth) - 1)
	} else {
		off = uint32(in)
	}

	pos = uint64(off) * entWidth

	fmt.Printf("Reading @ offset %d with position %d\n", off, pos)

	if i.size < pos+entWidth {
		return 0, 0, io.EOF
	}

	off = enc.Uint32(i.mmap[pos : pos+offWidth])
	pos = enc.Uint64(i.mmap[pos+offWidth : pos+entWidth])

	return off, pos, nil
}

func (i *index) Write(off uint32, pos uint64) error {
	if uint64(len(i.mmap)) < i.size+entWidth {
		return io.EOF
	}

	enc.PutUint32(i.mmap[i.size:i.size+offWidth], off)
	enc.PutUint64(i.mmap[i.size+offWidth:i.size+entWidth], pos)
	i.size += uint64(entWidth)

	return nil
}

func (i *index) Name() string {
	return i.file.Name()
}

func (i *index) Close() error {
	// ensure flushed to file
	if err := i.mmap.Sync(gommap.MS_SYNC); err != nil {
		return err
	}

	// ensure flushed to disk
	if err := i.file.Sync(); err != nil {
		return err
	}
	if err := i.file.Truncate(int64(i.size)); err != nil {
		return err
	}

	return i.file.Close()
}

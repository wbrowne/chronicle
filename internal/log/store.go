package log

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"sync"
)

var (
	enc = binary.BigEndian
)

const (
	lenWidth = 8 // record length (bytes)
)

type store struct {
	mu sync.Mutex
	*os.File
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*store, error) {
	fi, err := f.Stat()

	if err != nil {
		return nil, err
	}

	return &store{
		File: f,
		buf:  bufio.NewWriter(f),
		size: uint64(fi.Size()),
	}, nil
}

func (s *store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}
	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}

	bytesWritten := uint64(w) + lenWidth

	pos = s.size
	s.size += bytesWritten

	fmt.Printf("[w][store] Appending %d bytes of record (total %d) (0x%x) @ pos %d\n", w, bytesWritten, string(p), pos)
	fmt.Printf("Current store size = %d\n", s.size)

	return bytesWritten, pos, nil
}

func (s *store) ReadAt(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// flush to disk in case we haven't yet
	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	// find how many bytes we have to read to get the whole record
	size := make([]byte, lenWidth)
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}
	fmt.Printf("[r][store] Determining record length (with given pos %d): %d\n", pos, enc.Uint64(size))

	// fetch and return the record
	b := make([]byte, enc.Uint64(size))
	if _, err := s.File.ReadAt(b, int64(pos+lenWidth)); err != nil {
		return nil, err
	}
	fmt.Printf("[r][store] Reading record (%d bytes): 0x%x\n", len(b), string(b))

	return b, nil
}

func (s *store) Close() error {
	// ensure buffer is flushed
	if err := s.buf.Flush(); err != nil {
		return err
	}

	fmt.Println("Flushing store to disk")

	// ensure flushed to disk
	if err := s.File.Sync(); err != nil {
		return err
	}
	if err := s.File.Truncate(int64(s.size)); err != nil {
		return err
	}

	return s.File.Close()
}

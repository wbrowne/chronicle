package log

import (
	"fmt"
	"os"
	"path"

	"github.com/gogo/protobuf/proto"

	api "github.com/wbrowne/chronicle/api/v1"
)

type segment struct {
	store                  *store
	index                  *index
	baseOffset, nextOffset uint64
	config                 *Config
}

func newSegment(dir string, baseOffset uint64, c *Config) (*segment, error) {
	s := &segment{
		baseOffset: baseOffset,
		config:     c,
	}

	storeFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".store")),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		return nil, err
	}
	if s.store, err = newStore(storeFile); err != nil {
		return nil, err
	}

	indexFile, err := os.OpenFile(
		path.Join(dir, fmt.Sprintf("%d%s", baseOffset, ".index")),
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return nil, err
	}

	if s.index, err = newIndex(indexFile, c); err != nil {
		return nil, err
	}

	// if the index is empty, then the next record appended to the segment
	// would be the first record and its offset would be the segment’s base offset
	if off, _, err := s.index.Read(-1); err != nil {
		s.nextOffset = baseOffset
	} else {
		s.nextOffset = baseOffset + uint64(off) + 1
	}

	return s, nil
}

func (s *segment) Append(record *api.Record) (offset uint64, err error) {
	record.Offset = s.nextOffset

	p, err := proto.Marshal(record)
	if err != nil {
		return 0, err
	}

	fmt.Printf("[w][segment] Writing record: {value:\"%s\", offset:%d}\n", record.Value, record.Offset)

	// append to log
	_, pos, err := s.store.Append(p)
	if err != nil {
		return 0, err
	}

	// add index entry
	if err = s.index.Write(
		uint32(s.nextOffset-uint64(s.baseOffset)), // index offsets are relative to base offset
		pos,
	); err != nil {
		return 0, err
	}
	cur := s.nextOffset
	s.nextOffset++

	fmt.Printf("Segment base offset: %d, next offset: %d\n", s.baseOffset, s.nextOffset)

	return cur, nil
}

func (s *segment) Read(off uint64) (*api.Record, error) {
	_, pos, err := s.index.Read(int64(off - s.baseOffset))
	if err != nil {
		return nil, err
	}
	p, err := s.store.ReadAt(pos)
	if err != nil {
		return nil, err
	}
	record := &api.Record{}
	err = proto.Unmarshal(p, record)

	fmt.Printf("[r][segment] Returning record: {value:\"%s\", offset:%d}\n", record.Value, record.Offset)
	return record, err
}

func (s *segment) Close() error {
	if err := s.index.Close(); err != nil {
		return err
	}
	if err := s.store.Close(); err != nil {
		return err
	}
	return nil
}

func (s *segment) Remove() error {
	if err := s.index.Close(); err != nil {
		return err
	}

	if err := os.Remove(s.index.Name()); err != nil {
		return err
	}

	if err := os.Remove(s.store.Name()); err != nil {
		return err
	}

	return nil
}

func (s *segment) IsMaxed() bool {
	return s.store.size >= s.config.Segment.MaxStoreBytes ||
		s.index.size >= s.config.Segment.MaxIndexBytes
}

func nearestMultiple(j, k uint64) uint64 {
	if j >= 0 {
		return (j / k) * k
	}
	return ((j - k + 1) / k) * k
}

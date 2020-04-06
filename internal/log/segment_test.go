package log

import (
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"os"
	"testing"

	api "github.com/wbrowne/chronicle/api/v1"
)

func TestSegment(t *testing.T) {
	dir, _ := ioutil.TempDir("", "segment_test")
	defer os.RemoveAll(dir)

	want := &api.Record{Value: []byte("test")}

	c := &Config{
		Segment{
			MaxStoreBytes: 1024,
			MaxIndexBytes: entWidth * 3,
		},
	}

	s, err := newSegment(dir, 16, c)
	require.NoError(t, err)
	require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)
	require.False(t, s.IsMaxed())

	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		require.NoError(t, err)
		require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		require.NoError(t, err)
		// can't compare directly due to json field 'XXX_sizecache' being set
		require.Equal(t, want.Value, got.Value)
		require.Equal(t, want.Offset, got.Offset)
	}

	_, err = s.Append(want)
	require.Equal(t, io.EOF, err)

	require.True(t, s.IsMaxed()) // maxed index

	c = &Config{
		Segment{
			MaxStoreBytes: uint64(len(want.Value) * 3),
			MaxIndexBytes: 1024,
		},
	}
	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	require.True(t, s.IsMaxed()) // maxed store

	err = s.Remove()
	require.NoError(t, err)

	s, err = newSegment(dir, 16, c)
	require.NoError(t, err)

	require.False(t, s.IsMaxed())
}

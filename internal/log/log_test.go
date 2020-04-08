package log

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"

	api "github.com/wbrowne/chronicle/api/v1"
)

func TestLog(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T, log *Log){
		"append and read a record succeeds": testAppendRead,
		"offset out of range error":         testOutOfRangeErr,
		"init with existing segments":       testInitExisting,
	} {
		t.Run(scenario, func(t *testing.T) {
			dir, err := ioutil.TempDir("", "log_test")
			require.NoError(t, err)
			defer os.RemoveAll(dir)

			log, err := NewLog(dir, &Config{Segment{MaxStoreBytes: 32}})
			require.NoError(t, err)

			fn(t, log)
		})
	}
}

func testAppendRead(t *testing.T, log *Log) {
	record := &api.Record{
		Value: []byte("hello world"),
	}
	off, err := log.Append(record)
	require.NoError(t, err)
	require.Equal(t, uint64(0), off)

	read, err := log.Read(off)
	require.NoError(t, err)
	// can't compare directly due to json field 'XXX_sizecache' being set
	require.Equal(t, record.Value, read.Value)
	require.Equal(t, record.Offset, read.Offset)
}

func testOutOfRangeErr(t *testing.T, log *Log) {
	read, err := log.Read(1)
	require.Nil(t, read)

	require.Error(t, err)
}

func testInitExisting(t *testing.T, o *Log) {
	record := &api.Record{
		Value: []byte("hello world"),
	}
	recordEntries := 3
	for i := 0; i < recordEntries; i++ {
		_, err := o.Append(record)
		require.NoError(t, err)
	}
	require.NoError(t, o.Close())

	n, err := NewLog(o.Dir, o.Config)
	require.NoError(t, err)

	off, err := n.Append(record)
	require.NoError(t, err)
	require.Equal(t, uint64(recordEntries), off)
}

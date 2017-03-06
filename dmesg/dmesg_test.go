package dmesg

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseEntry(t *testing.T) {
	entry, err := parseEntry("6,0,0,-;Initializing cgroup subsys cpuset")
	require.NoError(t, err, "Expected no error")
	assert.Equal(t, &LogEntry{
		Level: LOG_INFO,
		Facility: LOG_KERN,
		SeqNum: 0,
		Timestamp: Timestamp(0),
		Message: "Initializing cgroup subsys cpuset",
		Tags: make(map[string]string),
	}, entry)

	entry, err = parseEntry("0,1,13,-;thonix: prebommand")
	require.NoError(t, err, "Expected no error")
	assert.Equal(t, &LogEntry{
		Level: LOG_EMERG,
		Facility: LOG_KERN,
		SeqNum: 1,
		Timestamp: Timestamp(13),
		Message: "thonix: prebommand",
		Tags: make(map[string]string),
	}, entry)
}

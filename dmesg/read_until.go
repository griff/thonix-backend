package dmesg

import (
	"time"
)

func ReadUntil(reader LogReadCloser, after time.Duration) ([]LogEntry, error) {
	ret := make([]LogEntry, 0)
	ch, errch := Channel(reader)
	var err error
	for {
		select {
		case entry := <-ch:
			ret = append(ret, entry)
		case err := <-errch:
			return ret, err
		case <-time.After(after):
			return ret, err
		}
	}
}

package dmesg

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DmesgLog struct {
	fd *os.File
}

func OpenDmesg() (*DmesgLog, error) {
	f, err := os.Open("/dev/kmsg")
	if err != nil {
		return nil, err
	}
	return &DmesgLog{
		fd: f,
	}, nil
}

func (l *DmesgLog) Close() error {
	return l.fd.Close()
}

func parseEntry(msg string) (*LogEntry, error) {
	split := strings.Split(msg, "\n")
	msg, tagsRaw := split[0], split[1:]
	tags := make(map[string]string)
	for _, tagRaw := range tagsRaw {
		if tagRaw == "" {
			continue
		}
		split := strings.SplitN(tagRaw, "=", 2)
		tags[split[0]] = split[1]
	}

	split = strings.SplitN(msg, ";", 2)
	optionsRaw, msg := split[0], split[1]

	options := strings.Split(optionsRaw, ",")

	faclev, err := strconv.ParseInt(options[0], 10, 32)
	if err != nil {
		return nil, err
	}
	lev := Level(faclev & 7)
	fac := Facility(faclev >> 3)

	seq, err := strconv.ParseInt(options[1], 10, 0)
	if err != nil {
		return nil, err
	}

	ts, err := strconv.ParseInt(options[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &LogEntry{
		Level:     lev,
		Facility:  fac,
		SeqNum:    int(seq),
		Timestamp: Timestamp(ts),
		Message:   msg,
		Tags:      tags,
	}, nil
}

func (l *DmesgLog) Read() (*LogEntry, error) {
	buf := make([]byte, 64000)
	n, err := l.fd.Read(buf)
	if err != nil {
		return nil, err
	}
	return parseEntry(string(buf[:n]))
}

func Write(msg string) (int, error) {
	f, err := os.OpenFile("/dev/kmsg", os.O_WRONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return fmt.Fprintln(f, msg)
}

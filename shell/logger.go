package shell

import (
	"bytes"
	"io"
	"sync"
)

type Logger struct {
	logs []byte
	mu   sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{
		logs: make([]byte, 0),
	}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	l.logs = append(l.logs, p...)
	l.mu.Unlock()
	return len(p), nil
}

func (l *Logger) Close() error {
	return nil
}

func (l *Logger) Reader() io.Reader {
	var b []byte
	l.mu.Lock()
	b = append([]byte(nil), l.logs...)
	l.mu.Unlock()
	return bytes.NewReader(b)
}

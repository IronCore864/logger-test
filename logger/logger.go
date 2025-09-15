package logger

import (
	"fmt"
	"io"
	"sync"
	"time"
)

const (
	timestampFormat = "2006-01-02T15:04:05.000Z07:00"
)

type Logger interface {
	Noticef(format string, v ...any)
}

var (
	logger     Logger
	loggerLock sync.Mutex
)

func Noticef(format string, v ...any) {
	loggerLock.Lock()
	defer loggerLock.Unlock()
	logger.Noticef(format, v...)
}

func SetLogger(l Logger) {
	loggerLock.Lock()
	logger = l
	loggerLock.Unlock()
}

// byteSliceWriter wraps a byte slice to implement io.Writer
type byteSliceWriter struct {
	buf *[]byte
}

func (w *byteSliceWriter) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}

type defaultLogger struct {
	w      io.Writer
	prefix string

	buf []byte
}

func (l *defaultLogger) Noticef(format string, v ...any) {
	l.buf = l.buf[:0]
	now := time.Now().UTC()
	l.buf = now.AppendFormat(l.buf, timestampFormat)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, l.prefix...)
	writer := &byteSliceWriter{buf: &l.buf}
	fmt.Fprintf(writer, format, v...)
	if len(l.buf) == 0 || l.buf[len(l.buf)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.w.Write(l.buf)
}

func New(w io.Writer, prefix string) Logger {
	return &defaultLogger{
		w:      w,
		prefix: prefix,
		buf:    make([]byte, 0, 256), // Preallocate some reasonable capacity.
	}
}

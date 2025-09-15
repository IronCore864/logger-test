package logger

import (
	"bytes"
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

type defaultLogger struct {
	w      io.Writer
	prefix string

	pool sync.Pool
}

func (l *defaultLogger) Noticef(format string, v ...any) {
	buf := l.pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer l.pool.Put(buf)
	now := time.Now().UTC()
	buf.WriteString(now.Format(timestampFormat))
	buf.WriteByte(' ')
	buf.WriteString(l.prefix)
	fmt.Fprintf(buf, format, v...)
	if buf.Len() == 0 || buf.Bytes()[buf.Len()-1] != '\n' {
		buf.WriteByte('\n')
	}
	l.w.Write(buf.Bytes())
}

func New(w io.Writer, prefix string) Logger {
	return &defaultLogger{
		w:      w,
		prefix: prefix,
		pool: sync.Pool{
			New: func() any {
				return &bytes.Buffer{}
			},
		},
	}
}

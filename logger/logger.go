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
	Notice(msg string)
}

var (
	logger     Logger
	loggerLock sync.Mutex
)

func Noticef(format string, v ...any) {
	loggerLock.Lock()
	defer loggerLock.Unlock()
	logger.Notice(fmt.Sprintf(format, v...))
}

func SetLogger(l Logger) {
	loggerLock.Lock()
	logger = l
	loggerLock.Unlock()
}

type defaultLogger struct {
	w      io.Writer
	prefix string

	buf []byte
}

func (l *defaultLogger) Notice(msg string) {
	l.buf = l.buf[:0]
	now := time.Now().UTC()
	l.buf = now.AppendFormat(l.buf, timestampFormat)
	l.buf = append(l.buf, ' ')
	l.buf = append(l.buf, l.prefix...)
	l.buf = append(l.buf, msg...)
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.w.Write(l.buf)
}

func New(w io.Writer, prefix string) Logger {
	return &defaultLogger{w: w, prefix: prefix}
}

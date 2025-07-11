package nekolog

import (
	"bytes"
	"runtime"
	"strings"
	"time"
)

type entry struct {
	l        *logger
	buffer   *bytes.Buffer
	m        map[string]interface{}
	level    Level
	time     time.Time
	file     string
	line     int
	function string
	format   string
	args     []interface{}
}

func NewEntry(l *logger) *entry {
	e := &entry{l: l, buffer: &bytes.Buffer{}, m: make(map[string]interface{}, 5)}

	return e
}

func (e *entry) write(level Level, format string, args ...interface{}) {
	if level < e.l.opt.level {
		return
	}
	e.level = level
	e.time = time.Now()
	e.format = format
	e.args = args

	if !e.l.opt.disableCaller {
		if pc, file, line, ok := runtime.Caller(2); !ok {
			e.file = "???"
			e.function = "???"
			e.line = 0
		} else {
			e.file, e.line, e.function = file, line, runtime.FuncForPC(pc).Name()
			e.function = e.function[strings.LastIndex(e.function, ".")+1:]
		}
	}

	e.fformat()
	e.wwrite()
	e.release()
}

func (e *entry) fformat() {
	_ = e.l.opt.formatter.Format(e)
}

func (e *entry) wwrite() {
	_, _ = e.l.opt.output.Write(e.buffer.Bytes())
}

func (e *entry) release() {
	e.buffer.Reset()
	e.l.entrypool.Put(e)
}

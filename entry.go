package nekolog

import (
	"bytes"
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
	e := &entry{l: l, buffer: &bytes.Buffer{}}

	return e
}

func (e *entry) write(level Level, format string, args ...interface{})

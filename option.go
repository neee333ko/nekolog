package nekolog

import (
	"io"
	"os"
)

const (
	FmtEmptySeparate = ""
)

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

type Level uint8

var levelMap = map[int]string{
	0: "Debug",
	1: "Info",
	2: "Warn",
	3: "Error",
	4: "Panic",
	5: "Fatal",
}

func levelUnmarshal(l Level) string {
	return levelMap[int(l)]
}

type option struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*option)

func initOptions(ops ...Option) *option {
	option := &option{}

	for _, op := range ops {
		op(option)
	}

	if option.output == nil {
		option.output = os.Stderr
	}

	if option.formatter == nil {
		option.formatter = &TextFormatter{}
	}

	return option
}

func WithOutput(w io.Writer) Option {
	return func(o *option) {
		o.output = w
	}
}

func WithLevel(l Level) Option {
	return func(o *option) {
		o.level = l
	}
}

func WithStdLevel(l Level) Option {
	return func(o *option) {
		o.stdLevel = l
	}
}

func WithFormatter(f Formatter) Option {
	return func(o *option) {
		o.formatter = f
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *option) {
		o.disableCaller = caller
	}
}

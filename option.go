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
		option.formatter = &JsonFormatter{}
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

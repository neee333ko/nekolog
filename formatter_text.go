package nekolog

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/term"
)

var levelColors = map[Level]string{
	DebugLevel: "\033[90m",
	InfoLevel:  "\033[32m",
	WarnLevel:  "\033[33m",
	ErrorLevel: "\033[31m",
	PanicLevel: "\033[35m",
	FatalLevel: "\033[1;31m",
}

type TextFormatter struct {
	disableBasicField bool
	disableColors     bool
	short             bool
}

func (f *TextFormatter) Format(e *entry) error {
	var str string

	if !f.disableBasicField {
		if f.short {
			e.file = e.file[strings.LastIndex(e.file, "/")+1:]
		}

		if !f.disableColors && isTerm(e.l.opt.output) {
			str = levelColors[e.level] + levelUnmarshal(e.level) + "\033[0m" + ": " + e.time.Format(time.RFC3339) + " " + e.file + ":" + strconv.Itoa(e.line) + " " + e.function + " "
		} else {
			str = levelUnmarshal(e.level) + ": " + e.time.Format(time.RFC3339) + " " + e.file + ":" + strconv.Itoa(e.line) + " " + e.function + " "
		}

	}

	if e.format == FmtEmptySeparate {
		str += fmt.Sprint(e.args...)
	} else {
		str += fmt.Sprintf(e.format, e.args...)
	}

	_, err := e.buffer.WriteString(str + "\n")

	return err
}

func isTerm(output io.Writer) bool {

	f, ok := output.(*os.File)
	if !ok {
		return false
	}

	switch f {
	case os.Stdout:
		if term.IsTerminal(int(os.Stdout.Fd())) {
			return true
		}
	case os.Stderr:
		if term.IsTerminal(int(os.Stderr.Fd())) {
			return true
		}
	}

	return false
}

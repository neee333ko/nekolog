package nekolog

import (
	"fmt"
	"strconv"
	"time"
)

type TextFormatter struct {
	disableBasicField bool
}

func (f *TextFormatter) Format(e *entry) error {
	var str string

	if !f.disableBasicField {
		str = levelUnmarshal(e.level) + ": " + e.time.Format(time.RFC3339) + " " + e.file + ":" + strconv.Itoa(e.line) + " " + e.function + " "
	}

	if e.format == FmtEmptySeparate {
		str += fmt.Sprint(e.args...)
	} else {
		str += fmt.Sprintf(e.format, e.args...)
	}

	_, err := e.buffer.WriteString(str + "\n")

	return err
}

package nekolog

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type JsonFormatter struct {
	disableBasicField bool
	short             bool
}

func (f *JsonFormatter) Format(e *entry) error {
	if !f.disableBasicField {
		e.m["level"] = levelUnmarshal(e.level)
		e.m["time"] = e.time.Format(time.RFC3339)

		if f.short {
			e.m["file"] = e.file[strings.LastIndex(e.file, "/")+1:] + ":" + strconv.Itoa(e.line)
		} else {
			e.m["file"] = e.file + ":" + strconv.Itoa(e.line)
		}
		e.m["func"] = e.function
	}

	if e.format == FmtEmptySeparate {
		e.m["msg"] = fmt.Sprint(e.args...)
	} else {
		e.m["msg"] = fmt.Sprintf(e.format, e.args...)
	}

	return jsoniter.NewEncoder(e.buffer).Encode(e.m)
}

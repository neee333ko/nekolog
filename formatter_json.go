package nekolog

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

type JsonFormatter struct {
	disableField bool
}

func (f *JsonFormatter) Format(e *entry) error {
	if !f.disableField {
		e.m["time"] = e.time.Format(time.RFC3339)
		e.m["level"] = e.level
		e.m["file"] = e.file
		e.m["line"] = e.line
		e.m["function"] = e.function
	}

	json := jsoniter.ConfigCompatibleWithStandardLibrary

	return json.NewEncoder(e.buffer).Encode(e.m)
}

package ui

import (
	"fmt"
	"io"
)

func safeGet(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func safeSet(str **string, value string) {
	if value == "" {
		*str = nil
	} else {
		*str = &value
	}
}
func emitScript(f io.Writer, cmd, path string, value interface{}) {
	fmt.Fprintf(f, "- command: %s\n", cmd)
	fmt.Fprintf(f, "  path: %s\n", path)
	if cmd != "delete" {
		if value == nil {
			fmt.Fprintf(f, "  value: null\n")
		} else {
			fmt.Fprintf(f, "  value: %#v\n", value)
		}
	}
}


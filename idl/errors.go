package idl

import (
	"fmt"
	"strings"

	"github.com/timtadh/lexmachine/machines"
)

func Errorf(m *machines.Match, format string, args ...interface{}) error {
	errInfo := fmt.Sprintf(format, args...)
	code := strings.Replace(string(m.Bytes), "\n", "\n\t", -1)
	prefix := fmt.Sprintf("syntax error: line:%d col:%d: %s\n%s", m.EndLine, m.EndColumn, errInfo, code)
	return fmt.Errorf(prefix)
}

package benchmark

import (
	"bytes"
	"fmt"
	"strings"
)

func Concat(start, add string) string {
	return start + add
}

func SprintfConcat(start, add string) string {
	return fmt.Sprintf("%s%s", start, add)
}

func BufferConcat(start, add string) string {
	var buffer bytes.Buffer
	buffer.WriteString(start)
	buffer.WriteString(add)
	return buffer.String()
}

func StringsConcat(start, add string) string {
	return strings.Join([]string{start, add}, "")
}

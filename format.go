package touzi

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

type Formatter interface {
	Format(input Result, format string) string
}

type DefaultFormatter struct{}

func prepareFormatString(format string) string {
	if format == "" {
		return "%d"
	}

	return "%" + format
}

func formatInt(format string, value int64) string {
	return fmt.Sprintf(prepareFormatString(format), value)
}

func formatUint(format string, value uint64) string {
	return fmt.Sprintf(prepareFormatString(format), value)
}

func formatBytes(format string, value []byte) string {
	if format == "" || format == "x" || format == "X" {
		formatted := hex.EncodeToString(value)

		if format == "X" {
			formatted = strings.ToUpper(formatted)
		}

		return formatted
	} else if format == "base64" {
		return base64.StdEncoding.EncodeToString(value)
	} else {
		return ""
	}
}

func (f *DefaultFormatter) Format(input Result, format string) string {
	switch input.(type) {
	case bool:
		if input.(bool) {
			return "true"
		}
		return "false"
	case int64:
		return formatInt(format, input.(int64))
	case uint64:
		return formatUint(format, input.(uint64))
	case []byte:
		return formatBytes(format, input.([]byte))
	}

	return ""
}

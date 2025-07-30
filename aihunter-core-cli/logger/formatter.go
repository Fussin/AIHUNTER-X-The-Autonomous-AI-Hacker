package logger

import (
	"encoding/json"
	"fmt"
)

// Formatter is an interface for formatting output.
type Formatter interface {
	Format(data interface{}) ([]byte, error)
}

// JSONFormatter is a formatter for JSON output.
type JSONFormatter struct{}

// Format formats the data as JSON.
func (f *JSONFormatter) Format(data interface{}) ([]byte, error) {
	return json.MarshalIndent(data, "", "  ")
}

// NewFormatter creates a new formatter based on the provided format.
func NewFormatter(format string) (Formatter, error) {
	switch format {
	case "json":
		return &JSONFormatter{}, nil
	default:
		return nil, fmt.Errorf("unknown format: %s", format)
	}
}

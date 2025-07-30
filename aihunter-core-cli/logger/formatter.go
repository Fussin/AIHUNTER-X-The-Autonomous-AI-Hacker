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

// ConsoleFormatter is a formatter for console output.
type ConsoleFormatter struct{}

// Format formats the data as a string.
func (f *ConsoleFormatter) Format(data interface{}) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", data)), nil
}

// NewFormatter creates a new formatter based on the provided format.
func NewFormatter(format string) (Formatter, error) {
	switch format {
	case "json":
		return &JSONFormatter{}, nil
	case "console":
		return &ConsoleFormatter{}, nil
	default:
		return nil, fmt.Errorf("unknown format: %s", format)
	}
}

package utils

import (
	"github.com/fatih/color"
)

// PrintError prints an error message in red.
func PrintError(err error) {
	color.Red("Error: %s", err.Error())
}

// PrintWarning prints a warning message in yellow.
func PrintWarning(msg string) {
	color.Yellow("Warning: %s", msg)
}

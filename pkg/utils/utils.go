// Package utils provides utility functions.
package utils

import (
	"fmt"
	"strings"
)

// Format formats a string with padding.
func Format(s string, padding int) string {
	pad := strings.Repeat(" ", padding)
	return fmt.Sprintf("%s%s", s, pad)
}

// Join joins strings with a separator.
func Join(sep string, items []string) string {
	return strings.Join(items, sep)
}

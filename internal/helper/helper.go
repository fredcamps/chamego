// Package helper provides internal helper functions.
package helper

import "strings"

// Sanitize sanitizes a string input.
func Sanitize(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

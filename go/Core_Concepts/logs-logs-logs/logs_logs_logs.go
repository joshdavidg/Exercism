package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {

		if c == '❗' {
			return "recommendation"
		}

		if c == '🔍' {
			return "search"
		}

		if c == '☀' {
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var sBuilder strings.Builder
	runeSlice := []rune(log)
	for i, r := range runeSlice {
		if r == oldRune {
			runeSlice[i] = newRune
		}
		fmt.Fprintf(&sBuilder, "%c", runeSlice[i])
	}
	return sBuilder.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

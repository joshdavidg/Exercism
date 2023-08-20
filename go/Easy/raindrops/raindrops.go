// Package raindrops implements functions to return a string of raindrop sounds.
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns a string of raindrop sounds based on rules that parse the number passed in.
// If the number does not fit the rules, the function will simply return the number as a string.
func Convert(number int) string {
	var s strings.Builder
	if number%3 == 0 {
		s.WriteString("Pling")
	}

	if number%5 == 0 {
		s.WriteString("Plang")
	}

	if number%7 == 0 {
		s.WriteString("Plong")
	}

	if s.Len() == 0 {
		return strconv.Itoa(number)
	}

	return s.String()
}

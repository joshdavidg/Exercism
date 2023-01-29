//Package twofer implements the function for the Two-fer go exercise.
package twofer

import "fmt"

// ShareWith function returns the string 'One for you, one for me.' if no name is given.
// If a name is given 'One for {name}, one for me.' is returned instead.
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return "One for you, one for me."
}

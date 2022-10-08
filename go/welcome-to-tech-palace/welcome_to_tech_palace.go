package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	var builder strings.Builder
	builder.WriteString("Welcome to the Tech Palace, ")
	builder.WriteString(strings.ToUpper(customer))
	return builder.String()
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var starBuilder strings.Builder
	var finalString strings.Builder
	for i := 0; i < numStarsPerLine; i++ {
		starBuilder.WriteString("*")
	}
	finalString.WriteString(starBuilder.String())
	finalString.WriteString("\n")
	finalString.WriteString(welcomeMsg)
	finalString.WriteString("\n")
	finalString.WriteString(starBuilder.String())
	return finalString.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(oldMsg, "*\n ")
}

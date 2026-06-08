package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	customer = strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + customer
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	star := strings.Repeat("*", numStarsPerLine)
	return star + "\n" + welcomeMsg + "\n" + star
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.TrimSpace(oldMsg)
	return oldMsg
}

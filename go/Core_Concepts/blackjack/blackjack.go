package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// Use map cause constant values and less code than large switch statement
	cardMap := map[string]int{
		"two": 2, "three": 3, "four": 4, "five": 5,
		"six": 6, "seven": 7, "eight": 8, "nine": 9,
		"ten": 10, "jack": 10, "queen": 10, "king": 10,
		"ace": 11,
	}
	val, keyFound := cardMap[card]

	if keyFound {
		return val
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	cardSum := ParseCard(card1) + ParseCard(card2)
	switch {
	case cardSum == 21 && ParseCard(dealerCard) < 10:
		return "W"
	case cardSum == 21 && ParseCard(dealerCard) >= 10:
		return "S"
	case cardSum == 22:
		return "P"
	case cardSum >= 17 && cardSum <= 20:
		return "S"
	case (cardSum >= 12 && cardSum <= 16) && ParseCard(dealerCard) < 7:
		return "S"
	default:
		return "H"
	}
}

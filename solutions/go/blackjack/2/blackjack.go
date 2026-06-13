package blackjack

var cardValue = map[string]int{
	"ace":   11,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"ten":   10,
	"jack":  10,
	"queen": 10,
	"king":  10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValue[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	vCard1 := ParseCard(card1)
	vCard2 := ParseCard(card2)
	vDealerCard := ParseCard(dealerCard)

	tCard := vCard1 + vCard2

	switch {
	case tCard <= 20 && tCard >= 17:
		return "S"

	case tCard <= 16 && tCard >= 12:
		if vDealerCard >= 7 {
			return "H"
		}
		return "S"

	// cards lower than 11
	case tCard <= 11:
		return "H"

	// pair of aces
	case vCard1 == 11 && vCard2 == 11:
		return "P"

	case tCard == 21:
		if vDealerCard == 11 || vDealerCard == 10 {
			return "S"
		}
		return "W"

	default:
		return "error"
	}
}

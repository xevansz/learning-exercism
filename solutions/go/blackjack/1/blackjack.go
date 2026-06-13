package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
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
		} else {
			return "S"
		}
	case tCard <= 11:
		return "H"
	case vCard1 == 11 && vCard2 == 11:
		return "P"
	case tCard == 21:
		if vDealerCard == 11 || vDealerCard == 10 {
			return "S"
		} else {
			return "W"
		}
	default:
		return "error"
	}
}

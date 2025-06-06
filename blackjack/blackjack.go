package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "king":
		fallthrough
	case "queen":
		fallthrough
	case "jack":
		fallthrough
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	switch {
	case sum > 21:
		return "P"
	case sum == 21 && dealerValue < 10:
		return "W"
	case sum >= 17:
		return "S"
	case sum <= 11:
		return "H"
	case dealerValue >= 7: // I know sum in [12, 16]
		return "H"
	default:
		return "S"
	}
}

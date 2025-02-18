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
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0

	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	isBlackjack := func(card1, card2 string) bool {
		return (card1 == "ace" && ParseCard(card2) == 10) || (card2 == "ace" && ParseCard(card1) == 10)
	}

	dealerShouldStand := func(dealerCard string) bool {
		dealerValue := ParseCard(dealerCard)
		return dealerValue == 11 || dealerValue == 10
	}

	if card1 == "ace" && card2 == "ace" {
		return "P"
	}

	if isBlackjack(card1, card2) {
		if !dealerShouldStand(dealerCard) {
			return "W"
		}
		return "S"
	}

	sum := ParseCard(card1) + ParseCard(card2)

	switch {
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16:
		dealerValue := ParseCard(dealerCard)
		if dealerValue >= 7 {
			return "H"
		}
		return "S"
	case sum <= 11:
		return "H"

	}
	return "S"
}

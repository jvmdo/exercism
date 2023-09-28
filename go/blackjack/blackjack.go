package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int

	switch card {
		case "two":
			value = 2
		case "three":
			value = 3
		case "four": 
			value = 4
		case "five":
			value = 5
		case "six":
			value = 6
		case "seven":
			value = 7
		case "eight":
			value = 8
		case "nine":
			value = 9
		// Alternatively,
		// case "ten", "jack", "queen", "king":
		case "ten":
			fallthrough
		case "jack":
			fallthrough
		case "queen":
			fallthrough
		case "king":
			value = 10
		case "ace":
			value = 11
		default:
			value = 0
	}

	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string

	cardsSum := ParseCard(card1) + ParseCard(card2)

	switch {
		case card1 == "ace" && card2 == "ace":
			decision = "P"

		case cardsSum == 21:
			if dealerCard == "ace" || ParseCard(dealerCard) == 10 {
				decision = "S"
			} else {
				decision = "W"
			}

		case cardsSum >= 17 && cardsSum < 21:
			decision = "S"

		case cardsSum >= 12 && cardsSum < 17:
			if ParseCard(dealerCard) >= 7 {
				decision = "H"
			} else {
				decision = "S"
			}

		case cardsSum < 12:
			decision = "H"

		default:
			decision = "W"
	}

	return decision
}

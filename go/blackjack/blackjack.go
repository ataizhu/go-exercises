package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return map[string]int{
		"ace": 11,
		"two": 2,
		"three": 3,
		"four": 4,	
		"five": 5,
		"six": 6,
		"seven": 7,	
		"eight": 8,
		"nine": 9,
		"ten": 10,
		"jack": 10,
		"queen": 10,
		"king": 10,
		"other": 0,
	}[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	handValue := ParseCard(card1) + ParseCard(card2)
	dealerValue := ParseCard(dealerCard)

	isPairOfAces := card1 == "ace" && card2 == "ace"
	hasBlackjack := handValue == 21
	dealerHasStrongCard := dealerValue >= 10 // 10, J, Q, K отдают 10; туз = 11

	switch {
	case isPairOfAces:
		return "P"
	case hasBlackjack && !dealerHasStrongCard:
		return "W"
	case hasBlackjack && dealerHasStrongCard:
		return "S"
	case handValue >= 17:
		return "S"
	case handValue >= 12 && handValue <= 16:
		if dealerValue >= 7 {
			return "H"
		}
		return "S"
	default: // handValue <= 11
		return "H"
	}
}

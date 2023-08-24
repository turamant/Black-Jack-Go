package card


var ranks = []string{"2", "3", "4", "5","6","7","8","9","10","J","Q","K","A"}
var suits = []rune{'\u2660','\u2661', '\u2662','\u2663'}

type Card struct {
	Rank string
	Suit string
}

type Hand struct {
	Cards []Card
}

type Deck struct {
	Cards []Card
}


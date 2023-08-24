package card

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)


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

func (h Hand) Value() int{
	score := 0
	haveAce := 0
	for index:=0;index < len(h.Cards);index++{
		if h.Cards[index].Rank != "A" && h.Cards[index].Rank != "K" && h.Cards[index].Rank != "Q" &&
		   h.Cards[index].Rank != "J"{
			intRank, _ := strconv.Atoi(h.Cards[index].Rank)
			score += intRank
		} else if h.Cards[index].Rank == "K" || h.Cards[index].Rank == "Q" || h.Cards[index].Rank == "J"{
			score += 10
		} else if h.Cards[index].Rank == "A"{
			score += 11
			haveAce += 1
		}
	}
	if score > 21 && haveAce > 1{
		score -= 10 * haveAce
	} 
	return score
}

func (h *Hand) AddCard(card Card){
	h.Cards = append(h.Cards, card)
}

func (h *Hand) Display(){
	for _,card := range h.Cards{
		fmt.Print(card.Rank + card.Suit + "..")
	}
}


func (deck *Deck) InitializeDeck() Deck{
	for _, suit := range suits {
		for _, rank := range ranks{
			deck.Cards = append(deck.Cards, Card{Rank: rank, Suit: string(suit)})
		} 
	}
	deck.Shuffle()
	return *deck
}


func (deck *Deck) Shuffle(){
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(deck.Cards), func(i, j int) {deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]})
}

func (deck *Deck) DealCard() Card{
	card := deck.Cards[0]
	deck.Cards = deck.Cards[1:]
	return card
}

func (deck *Deck) Display(){
	for _, card := range deck.Cards{
		fmt.Print(card.Rank + card.Suit + " | ")
	}
}

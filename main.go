package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/turamant/blackjack/card"
)
func main()  {
	gameOver := false
	myDeck := card.Deck{}
	myDeck.InitializeDeck()

	kazinoHand := card.Hand{}
	playerHand := card.Hand{}

	for i:=1;i<=2;i++{
		card := myDeck.DealCard()
		kazinoHand.AddCard(card)
		card = myDeck.DealCard()
		playerHand.AddCard(card)
	}
	playerHand.Display()
	fmt.Println("Ещё карту ? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := reader.ReadRune()
	for {
		if res != 'y'{
			break
		}
		card := myDeck.DealCard()
		playerHand.AddCard(card)
		playerHand.Display()
		if playerHand.Value() >21 {
			fmt.Println("У вас больше 21. Вы проиграли!")
			gameOver = true
			break
		}
		fmt.Println("Ещё карту ? (y/n)")
		reader := bufio.NewReader(os.Stdin)
		res, _, _ = reader.ReadRune()
	}
	if !gameOver {
		for {
			if kazinoHand.Value() >21 {
				fmt.Println("Казино больше 21. Вы выиграли")
				gameOver = true
				break
			}
			if kazinoHand.Value() < 17{
				card := myDeck.DealCard()
				kazinoHand.AddCard(card)
			} else {
				break
			}
		}
	}
	if !gameOver {
		if playerHand.Value() > kazinoHand.Value(){
			fmt.Println("Вы выиграли!")
		}else if playerHand.Value() == kazinoHand.Value(){
			fmt.Println("Ничья")
		}else{
			fmt.Println("Казино победило!")
		}
	}





}
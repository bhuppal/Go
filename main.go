package main

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards := newDeck()
	//cards.saveToFile("my_cards")


	/*
		fmt.Println(cards.toString())
		cardDetails := []byte(cards.toString())
		err := ioutil.WriteFile("Cards.txt", cardDetails, 0644)

		if err != nil {
			log.Fatal(err)
		}
	*/
/*
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	fmt.Println("******Card******")

	*/
}

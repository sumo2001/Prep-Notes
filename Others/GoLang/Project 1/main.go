package main

import "fmt"

func main() {
	fmt.Println("NewDeck Function:")
	cards := newDeck()
	cards.print()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Deal Function:")
	givenCards, LeftOverCards := deal(cards, 7)
	givenCards.print()
	fmt.Println("*******************************************")
	LeftOverCards.print()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Shuffle Function:")
	cards.shuffle()
	cards.print()
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("String to Byte Conversion:")
	out := cards.toString()
	fmt.Println(out)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Saved to a newfile")
	cards.saveToFile("newfile")
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Read from newfile")
	newRead := readFromFile("newfile")
	newRead.print()

}

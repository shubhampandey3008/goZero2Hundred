package main

import "fmt"

func main() {
	var cards string = "Five of Hearts"
	fmt.Println(cards)

	// The line below is the same as the above
	cards1 := "Nine of Spades"
	fmt.Println(cards1)

	// No need to use := when we reassign the value to cards1
	cards1 = "King of Hearts"
}

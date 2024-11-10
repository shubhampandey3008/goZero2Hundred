package main

import "fmt"

func main() {

	store := map[string]int{
		"Biscuit":  10,
		"Chips":    10,
		"Battery":  20,
		"Keyboard": 500,
	}

	// This is the first Print Statement
	fmt.Println(store)

	// Another way of defining a Map
	var nextStore map[string]int
	fmt.Printf(" The value of nextStore is %v \n", nextStore)

	// Adding elements to map
	store["Banana"] = 70

	// This is the second Print Statement
	fmt.Println(store)

	// Delete an element from map
	delete(store, "Chips")

	// Final Print statement
	fmt.Println(store)

	// unstructuring map
	for prod, cost := range store {
		fmt.Println(prod, " cost is ", cost)
	}
}

package main

import "fmt"

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

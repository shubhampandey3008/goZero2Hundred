package main

import (
	"fmt"
	"io"
	"os"
)

type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

type Square struct {
	sideLength float64
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {
	var newTriangle Triangle
	newTriangle.base = 4
	newTriangle.height = 10

	printArea(newTriangle)

	// Printing the content of a file in hard drive to the cl . The name of the file is passed
	// as command line arguement in the terminal

	// Step 1 : Write a text to file
	filetext := "This is the content of the file"

	// Step 2 : Taking the filename from command line
	filename := os.Args[1]

	// Step 3 : Creating a new file and writing the content to it
	file, _ := os.OpenFile(filename, os.O_CREATE, 0666)
	os.WriteFile(filename, []byte(filetext), 0666)

	// Step 4 : Now , we will read the file and print to the commandline
	fileBytes, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading the file")
		os.Exit(1)
	}

	fmt.Println(string(fileBytes))

	// Since the file struct implements the reader interface , we can directly use this
	io.Copy(os.Stdout, file)
}

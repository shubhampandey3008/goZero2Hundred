package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error getting the site data. Exiting.............")
		os.Exit(1)
	}

	bodySlice := make([]byte, 999999)

	resp.Body.Read(bodySlice)

	// Writing it to a file
	file, err := os.Create("google.html")
	if err != nil {
		fmt.Println("Error Opening the file")
		os.Exit(1)
	}
	defer file.Close()
	file.Write(bodySlice)
	fmt.Println("Wrote to file successfully!!!!")

	// fmt.Printf("Body from the site : %v", string(bodySlice))
}

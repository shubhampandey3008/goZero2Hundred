package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://linkedin.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkIfUp(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkIfUp(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		c <- link + " is not working"
		return
	}

	c <- link + " is up and running"
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	// Creating a channel
	ch := make(chan string)

	links := []string{
		"http://google.com",
		"http://linkedin.com",
		"http://youtube.com",
		"http://facebook.com",
	}

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, ch)
		}(l)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is not working!!!!!")
		ch <- link
		return
	}
	fmt.Println(link, "is up and running!!!")
	ch <- link

}

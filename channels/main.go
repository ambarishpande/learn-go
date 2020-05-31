package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://womenscriczone.com",
		"http://golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// Function Literal. Same as Lambda. An Unnamed function.
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(s string, c chan string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println(s, "might be down!")
		c <- s
		return
	}
	fmt.Println(s, "is up!")
	c <- s
}

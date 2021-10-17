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
		"http://twitter.com",
		"http://amazon.com",
	}

	// Creating a Channel for Communication
	c := make(chan string)

	// Creating Go Routines for all links
	for _, link := range links {
		go checkLink(link, c)
	}

	// Using Anonymous Functions
	// ! Don't let Variable Capture
	for l := range c {
		go func(l string) {
			time.Sleep(3 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link, "seems to be OK!")
	}

	c <- link
}

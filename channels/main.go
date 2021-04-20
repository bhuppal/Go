package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	siteList := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}
	c := make(chan string)
	for _, link := range siteList {
		go checkLink(link, c)
	}

	/*
		for i:=0; i < len(siteList); i++{
		fmt.Println(<-c)
		}
	*/

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(" might be down!")
		c <- link
		return
	}

	fmt.Println(link)
	c <- link
}

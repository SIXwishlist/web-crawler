package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide a website url")
		return
	}

	rootUrl := os.Args[1]

	visitedLinks := make(map[string]bool)
	jobs := make(chan string)
	foundLinks := make(chan string)
	done := make(chan bool)

	go func() { foundLinks <- rootUrl }()

	
	

	fmt.Println("New link: ", link)
	fmt.Println(visitedLinks)
	if !visitedLinks[link] {
		visitedLinks[link] = true
		unvisitedLinksCount++
		go func(link string) {
			newLinks := crawl(link)
			for i := 0; i < len(newLinks); i++ {
				newLink := newLinks[i]
				links <- newLink
			}
		}(link)
	}
}

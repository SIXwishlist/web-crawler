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
	links := make(chan string, 100)

	go func() { links <- rootUrl }()

	for link := range links {
		fmt.Println("Link to visit: ", link)
		if !visitedLinks[link] {
			visitedLinks[link] = true
			go func(link string) {
				newLinks := crawl(link)
				for i := 0; i < len(newLinks); i++ {
					newLink := newLinks[i]
					fmt.Println("newLink: ", newLink)
				}
			}(link)
		}
	}
}

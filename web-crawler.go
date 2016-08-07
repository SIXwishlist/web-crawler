package main

import (
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	for i:= 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

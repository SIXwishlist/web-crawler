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
			for {
				link, ok := <-unseenLinks
				if ok {
					foundLinks := crawl(link)
					worklist <- foundLinks
				} else {
					return
				}
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for {
		list, ok := <-worklist
		if ok {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		} else {
			close(worklist)
			close(unseenLinks)
		}
	}
}

package main

import (
	"flag"
	"fmt"
)

func main() {
	foundLinks := make(chan []string)
	unseenLinks := make(chan string)

	workers := flag.Int("w", 1, "Number of concurrent workers to perform requests")
	url := flag.String("u", "", "Starting URL")
	flag.Parse()

	go func() { foundLinks <- []string{*url} }()

	fmt.Println("No of workers:", *workers)
	for i:= 0; i < *workers; i++ {
		go func() {
			for link := range unseenLinks {
				crawl(i, link, foundLinks)
			}
		}()
	}

	// Crawl the web concurrently.
	seen := make(map[string]bool)

	for n := 1; n > 0; n-- {
		fmt.Println("n", n)
		list := <-foundLinks
		fmt.Println("Got list", list)
		for _, link := range list {
			fmt.Println("Checking link", link)
			fmt.Println("Seen is", seen)
			if !seen[link] {
				seen[link] = true
				n++
				fmt.Println("Adding", link)
				unseenLinks <- link
			}
		}
	}
}

package main

import (
	"flag"
)

var (
	workers = flag.Int("w", 1, "Number of concurrent workers to perform requests")
	startingUrl = flag.String("u", "", "Starting URL")
)

func main() {
	var (
		foundLinks = make(chan []string)
		unseenLinks = make(chan string)
		seen = make(map[string]bool)
	)

	flag.Parse()

	go func() { foundLinks <- []string{*startingUrl} }()

	for i:= 0; i < *workers; i++ {
		go worker(i, unseenLinks, foundLinks)
	}

	for n := 1; n > 0; n-- {
		list := <-foundLinks
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				unseenLinks <- link
			}
		}
	}
}

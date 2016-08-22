package main

import (
	"flag"
	"os"
)

var (
	foundLinks  = make(chan []string)
	unseenLinks = make(chan string)
	results     = make(chan pageInfo)
	seen        = make(map[string]bool)
	output			= os.Stdout
	workers     = flag.Int("w", 1, "Number of concurrent workers to perform requests")
	startingUrl = flag.String("u", "", "Starting URL")
)

func main() {
	flag.Parse()

	for i := 0; i < *workers; i++ {
		worker := NewWorker()
		go worker.Start(i, unseenLinks, foundLinks, results)
	}

	go Printer(output, results)

	go func() { foundLinks <- []string{*startingUrl} }()
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

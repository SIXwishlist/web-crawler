package main

import (
	"flag"
)

func startWorkers(workers int, newWorker WorkerConstructor, unseenLinks <-chan string, foundLinks chan<- []string) {
	for i:= 0; i < workers; i++ {
		worker := newWorker()
		go worker.Start(i, unseenLinks, foundLinks)
	}
}

func dispatchLinks(startingUrl string, foundLinks chan []string, unseenLinks chan<- string, seen map[string]bool) {
	go func() { foundLinks <- []string{startingUrl} }()
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

	startWorkers(*workers, NewWorker, unseenLinks, foundLinks)
	dispatchLinks(*startingUrl, foundLinks, unseenLinks, seen)
}

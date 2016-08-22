package main

import (
	"flag"
	"io"
	"os"
)

func startWorkers(workers int, newWorker WorkerConstructor, unseenLinks <-chan string, foundLinks chan<- []string, results chan<- pageInfo) {
	for i := 0; i < workers; i++ {
		worker := newWorker()
		go worker.Start(i, unseenLinks, foundLinks, results)
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

func startPrinter(output io.Writer, results <-chan pageInfo) {
	go Printer(output, results)
}

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

	startWorkers(*workers, NewWorker, unseenLinks, foundLinks, results)
	startPrinter(output, results)
	dispatchLinks(*startingUrl, foundLinks, unseenLinks, seen)
}

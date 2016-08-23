package main

import (
	"flag"
	"os"
	"io"
)

var (
	unseenLinks = make(chan string)
	found       = make(chan pageInfo)
	seen        = make(map[string]bool)
	output      = os.Stdout
	workers     = flag.Int("w", 1, "Number of concurrent workers to perform requests")
	startingUrl = flag.String("u", "", "Starting URL")
)

func print(output io.Writer, page pageInfo) {
	io.WriteString(output, "Page: "+page.page+"\n")

	io.WriteString(output, "  Links:\n")
	for _, link := range page.links {
		io.WriteString(output, "    - "+link+"\n")
	}

	io.WriteString(output, "  Assets:\n")
	for _, asset := range page.assets {
		io.WriteString(output, "    - "+asset+"\n")
	}
}

func main() {
	flag.Parse()

	for i := 0; i < *workers; i++ {
		worker := NewWorker()
		go worker.Start(i, unseenLinks, found)
	}

	go func() { found <- pageInfo{links: []string{*startingUrl}} }()
	for n := 1; n > 0; n-- {
		page := <-found
		print(output, page)
		for _, link := range page.links {
			if !seen[link] {
				seen[link] = true
				n++
				unseenLinks <- link
			}
		}
	}
}

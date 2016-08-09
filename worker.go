package main

import(
	"fmt"
)

func worker(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for link := range unseenLinks {
		fmt.Println("Worker #", id, "crawling", link)

		links := extractLinks(link, fetcher{})

		go func() { foundLinks <- links }()
	}
}

func extractLinks(link string, fetcher Fetcher) []string {
	body := fetcher.Fetch(link)
	links := body.ExtractLinks()
	return links
}

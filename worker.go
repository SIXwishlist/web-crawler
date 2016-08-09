package main

import(
	"fmt"
	"net/http"
)

func worker(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for link := range unseenLinks {
		fmt.Println("Worker #", id, "crawling", link)

		links := extractLinks(link, fetcher{client: &http.Client{}})

		go func() { foundLinks <- links }()
	}
}

func extractLinks(link string, fetcher Fetcher) (links []string) {
	doc, err := fetcher.Fetch(link)
	if err != nil {
		return
	}
	links = doc.ExtractLinks()
	return
}

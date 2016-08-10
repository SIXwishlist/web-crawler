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

func extractLinks(link string, fetcher Fetcher) (links []string) {
	body, err := fetcher.Fetch(link)
	if err != nil {
		return
	}
	fmt.Println("Fetched page", link)
	doc := htmlDoc{body: body}
	links = doc.ExtractLinks()
	return
}

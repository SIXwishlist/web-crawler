package main

import(
	"fmt"
)

func worker(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for link := range unseenLinks {
		links := extractLinks(link, fetcher{})

		go func() { foundLinks <- links }()
	}
}

func extractLinks(link string, fetcher Fetcher) (links []string) {
	body, err := fetcher.Fetch(link)
	if err != nil {
		return
	}
	fmt.Println("Page:", link)
	doc := NewHtmlDoc(body, link)
	links = doc.ExtractInternalLinks()
	printLinks(links)
	return
}

func printLinks(links []string) {
	fmt.Println("  Links:")
	for _, link := range links {
		fmt.Println("    -", link)
	}
}

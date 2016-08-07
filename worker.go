package main

import(
	"fmt"
)

func worker(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for link := range unseenLinks {
		crawl(id, link, foundLinks)
	}
}

func crawl(id int, link string, foundLinks chan<- []string) {
	fmt.Println("Worker #", id, "crawling", link)
	links := extractLinks(link)
	go func() { foundLinks <- links }()
}

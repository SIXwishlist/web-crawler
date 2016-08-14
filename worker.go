package main

import (
	"fmt"
)

type Worker struct {
	fetcher    Fetcher
	newHtmlDoc HtmlDocConstructor
}

type WorkerConstructor func() *Worker

func NewWorker() *Worker {
	return &Worker{fetcher: fetcher{}, newHtmlDoc: NewHtmlDoc}
}

func (this Worker) Start(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for link := range unseenLinks {
		fmt.Println("Worker ", id, "crawling", link)
		links := this.extractLinks(link)

		go func() { foundLinks <- links }()
	}
}

func (this Worker) extractLinks(link string) (links []string) {
	body, err := this.fetcher.Fetch(link)
	if err != nil {
		return
	}
	doc := this.newHtmlDoc(body, link)
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

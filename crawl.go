package main

import(
	"fmt"
)

func crawl(id int, link string, worklist chan<- []string) {
	fmt.Println("Worker #", id, "crawling", link)
	links := make([]string, 3)
	links[0] = "Link1"
	links[1] = "Link2"
	links[2] = "Link3"
	fmt.Println("Worker #", id, "found", links)
	worklist <- links
}

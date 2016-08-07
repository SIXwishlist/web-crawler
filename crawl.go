package main

import(
	"fmt"
)

func crawl(link string, worklist chan<- []string) {
	fmt.Println("Crawling ", link)
	links := make([]string, 2)
	links[0] = "foo"
	links[1] = "bar"
	worklist <- links
}

package main

import(
	"fmt"
)

func crawl(link string) []string {
	fmt.Println("Crawling ", link)
	links := make([]string, 1)
	links[0] = "foo"
	return links
}

package main

import "testing"

func testWorker(id int, unseenLinks <-chan string, foundLinks chan<- []string) {
	for _ = range unseenLinks {
		links := make([]string, 3)
		links[0] = "Link1"
		links[1] = "Link2"
		links[2] = "Link3"
		go func() { foundLinks <- links }()
	}
}

func TestWebCrawler(t *testing.T) {
	var (
		startingUrl = "http://tomblomfield.com/"
		workers = 2
		foundLinks = make(chan []string)
		unseenLinks = make(chan string)
		seen = make(map[string]bool)
	)

	startWorkers(workers, testWorker, unseenLinks, foundLinks)
	dispatchLinks(startingUrl, foundLinks, unseenLinks, seen)

	links := []string{startingUrl, "Link1", "Link2", "Link3"}
	for _, link := range links {
		found := false

		for seenLink, _ := range seen {
			if link == seenLink {
				found = true
			}
		}

		if !found {
			t.Error(link, "has not been crawled")
		}
	}
}

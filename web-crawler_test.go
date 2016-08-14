package main

import "testing"

type testFetcher struct {}

func (this testFetcher) Fetch(url string) (string, error) {
	return "", nil
}

type testHtmlDoc struct {}

func newTestHtmlDoc(body string, address string) HtmlDoc {
	return testHtmlDoc{}
}

func (this testHtmlDoc) ExtractInternalLinks() []string {
	return []string{"Link1","Link2", "Link3"}
}

func newTestWorker() *Worker {
	return &Worker{fetcher: testFetcher{}, newHtmlDoc: newTestHtmlDoc}
}

func TestWebCrawler(t *testing.T) {
	var (
		startingUrl = "http://tomblomfield.com/"
		workers = 2
		foundLinks = make(chan []string)
		unseenLinks = make(chan string)
		seen = make(map[string]bool)
	)

	startWorkers(workers, newTestWorker, unseenLinks, foundLinks)
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

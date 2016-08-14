package main

import "testing"

type testFetcher struct {}

func (this testFetcher) Fetch(url string) (string, error) {
	return `
<html>
	<body>
		<a href='http://tomblomfield.com/1'>1</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://google.com'>Google</a>
		<a href='http://google.com/'>Google</a>
		<a href='/page3'>3</a>
	</body>
</html>`, nil
}

func newTestWorker() *Worker {
	return &Worker{fetcher: testFetcher{}, newHtmlDoc: NewHtmlDoc}
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

	links := []string{"http://tomblomfield.com/1","http://tomblomfield.com/2", "http://tomblomfield.com/page3"}
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

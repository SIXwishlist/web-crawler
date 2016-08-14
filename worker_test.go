package main

import "testing"

type testFetcherWorker struct{}

func (this testFetcherWorker) Fetch(url string) (string, error) {
	return "", nil
}

type testHtmlDoc struct{}

func newTestHtmlDoc(body string, address string) HtmlDoc {
	return testHtmlDoc{}
}

func (this testHtmlDoc) ExtractInternalLinks() []string {
	return []string{"Link1", "Link2", "Link3"}
}

func TestExtractLinks(t *testing.T) {
	worker := Worker{fetcher: testFetcherWorker{}, newHtmlDoc: newTestHtmlDoc}

	expectedLinks := []string{"Link1", "Link2", "Link3"}
	links := worker.extractLinks("http://tomblomfield.com")

	if !equalStringSlices(links, expectedLinks) {
		t.Error("Expected links:", expectedLinks, "actual links", links)
	}
}

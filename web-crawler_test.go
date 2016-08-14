package main

import(
	"testing"
	"strings"
)

type testFetcher struct {}

func (this testFetcher) Fetch(url string) (string, error) {
	return `
<html>
	<head>
		<script type="text/javascript" async="" src="http://www.google-analytics.com/ga.js"></script>
	</head>
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

type testOutput struct {
}

func (this testOutput) Write(p []byte) (n int, err error) {
	return
}

var receivedOutput = ""

func (this testOutput) WriteString(s string) (n int, err error) {
	receivedOutput = receivedOutput + s
	return
}

func TestWebCrawler(t *testing.T) {
	var (
		startingUrl = "http://tomblomfield.com/"
		workers = 2
		foundLinks = make(chan []string)
		unseenLinks = make(chan string)
		results = make(chan pageResult)
		seen = make(map[string]bool)
		output = testOutput{}
	)

	startWorkers(workers, newTestWorker, unseenLinks, foundLinks, results)
	startPrinter(output, results)
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

	expectedOutput := `
Page: http://tomblomfield.com
  Assets:
    - http://www.google-analytics.com/ga.js
  Links:
    - http://tomblomfield.com/1
    - http://tomblomfield.com/2
    - http://tomblomfield.com/page3`

	if !strings.Contains(receivedOutput, expectedOutput) {
		t.Error("\n" + receivedOutput + "\n", "doesn't contain:", "\n" + expectedOutput)
	}
}

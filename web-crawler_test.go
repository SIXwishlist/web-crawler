package main

import "testing"

func TestWebCrawler(t *testing.T) {
	*workers = 2
	*startingUrl = "http://tomblomfield.com/"
	// test that crawler spawns 2 workers
	// each worker returns 3 links
	// all links are crawled
	// and program terminates with 0
	main()
}

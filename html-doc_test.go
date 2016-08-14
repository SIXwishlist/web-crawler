package main

import "testing"

var (
	bodyHtml = `
<html>
	<body>
		<a href='http://tomblomfield.com/1'>1</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://google.com'>Google</a>
		<a href='/page3'>3</a>
	</body>
</html>`
)

func equalStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		found := false

		for j := range s2 {
			if s1[i] == s2[j] {
				found = true
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func TestExtractInternalLinks(t *testing.T) {
	doc := htmlDoc{body: bodyHtml, domain: "http://tomblomfield.com"}
	links := doc.ExtractInternalLinks()
	expectedLinks := []string{"http://tomblomfield.com/1","http://tomblomfield.com/2", "http://tomblomfield.com/page3"}

	if !equalStringSlices(expectedLinks, links) {
		t.Error("Expected links:", expectedLinks, "actual links", links)
	}
}

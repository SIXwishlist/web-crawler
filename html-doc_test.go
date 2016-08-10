package main

import "testing"

var (
	bodyHtml = `
<html>
	<body>
		<a href='http://tomblomfield.com/1'>1</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://google.com'>Google</a>
	</body>
</html>`
)

func equalStringSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func TestExtractLinks(t *testing.T) {
	doc := htmlDoc{body: bodyHtml}
	links := doc.ExtractLinks()
	expectedLinks := []string{"http://tomblomfield.com/1","http://tomblomfield.com/2","http://google.com"}

	if !equalStringSlices(expectedLinks, links) {
		t.Error("Expected links:", expectedLinks, "actual links", links)
	}
}

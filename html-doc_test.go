package main

import "testing"

var (
	bodyHtml = `
<html>
	<head>
		<script type="text/javascript" async="" src="http://www.google-analytics.com/ga.js"></script>
	</head>
	<body>
    <img src="/puppy.jpeg" />
		<a href='http://tomblomfield.com/1'>1</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://tomblomfield.com/2'>2</a>
		<a href='http://google.com'>Google</a>
		<a href='http://google.com/'>Google</a>
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

func TestExtractPageInfo(t *testing.T) {
	doc := htmlDoc{body: bodyHtml, domain: "http://tomblomfield.com", address: "http://tomblomfield.com"}
	info := doc.ExtractPageInfo()
	expectedPageInfo := pageInfo{page: "http://tomblomfield.com", links: []string{"http://tomblomfield.com/1","http://tomblomfield.com/2", "http://tomblomfield.com/page3"}, assets: []string{"http://www.google-analytics.com/ga.js", "/puppy.jpeg"}}


	if !equalStringSlices(expectedPageInfo.links, info.links) {
		t.Error("Expected links:", expectedPageInfo.links, "actual links", info.links)
	}

	if !equalStringSlices(expectedPageInfo.assets, info.assets) {
		t.Error("Expected assets:", expectedPageInfo.assets, "actual assets", info.assets)
	}

	if expectedPageInfo.page != info.page {
		t.Error("Expected page:", expectedPageInfo.page, "actual page", info.page)
	}
}

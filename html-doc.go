package main

type HtmlDoc interface {
	ExtractLinks() []string
}

type htmlDoc struct {
	body string
}

func (document htmlDoc) ExtractLinks() []string {
	var links []string
	return links
}

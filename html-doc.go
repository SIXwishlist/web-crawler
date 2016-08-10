package main

type HtmlDoc interface {
	ExtractLinks() []string
	ReadBody() string
}

type htmlDoc struct {
	body string
}

func (document htmlDoc) ReadBody() string {
	return "<html></html>"
}

func (document htmlDoc) ExtractLinks() []string {
	var links []string
	return links
}

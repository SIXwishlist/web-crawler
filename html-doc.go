package main

import "io"

type HtmlDoc interface {
	ExtractLinks() []string
	ReadBody() string
}

type htmlDoc struct {
	body io.ReadCloser
}

func (document htmlDoc) ReadBody() string {
	return "<html></html>"
}

func (document htmlDoc) ExtractLinks() []string {
	var links []string
	return links
}

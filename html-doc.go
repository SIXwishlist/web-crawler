package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

type HtmlDoc interface {
	ExtractLinks() []string
	ReadBody() string
}

type htmlDoc struct {
	body string
}

func (this htmlDoc) ReadBody() string {
	return "<html></html>"
}

func selectLinks(n *html.Node, buf []string) []string {
	fmt.Println(n.Data)
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				fmt.Println(attr.Val)
				buf = append(buf, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buf = selectLinks(c, buf)
	}
	return buf
}

func (this htmlDoc) ExtractLinks() []string {
	var links []string

	doc, err := html.Parse(strings.NewReader(this.body))
	if err != nil {
		return nil
	}

	links = selectLinks(doc, links)

	return links
}

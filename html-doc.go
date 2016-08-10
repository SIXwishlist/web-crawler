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

func (this htmlDoc) ExtractLinks() []string {
	var links []string

	doc, err := html.Parse(strings.NewReader(this.body))
	if err != nil {
		return nil
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
					links = append(links, attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

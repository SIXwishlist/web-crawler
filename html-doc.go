package main

import (
	"golang.org/x/net/html"
	"strings"
	"regexp"
)

var domainRegex = regexp.MustCompile(`https?:\/\/([\w\d])+(\.\w+)*`)

type HtmlDoc interface {
	ExtractInternalLinks() []string
	ReadBody() string
}

type htmlDoc struct {
	body string
	domain string
}

func NewHtmlDoc(body string, address string) *htmlDoc {
	domain := domainRegex.FindString(address)

	return &htmlDoc{body: body, domain: domain}
}

func (this htmlDoc) ReadBody() string {
	return "<html></html>"
}

func selectLinks(n *html.Node, buf []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				buf = append(buf, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buf = selectLinks(c, buf)
	}
	return buf
}

func filterInternalLinks(links []string, domain string) (internalLinks []string) {
	uniqueInternalLinks := make(map[string]bool)

	for _, link := range links {
		if strings.HasPrefix(link, domain) {
			uniqueInternalLinks[link] = true
		} else if strings.HasPrefix(link, "/") {
			uniqueInternalLinks[domain+link] = true
		}
	}

	for link, _ := range uniqueInternalLinks {
		internalLinks = append(internalLinks, link)
	}

	return
}

func (this htmlDoc) ExtractInternalLinks() []string {
	var links []string

	doc, err := html.Parse(strings.NewReader(this.body))
	if err != nil {
		return nil
	}

	links = selectLinks(doc, links)
	links = filterInternalLinks(links, this.domain)

	return links
}

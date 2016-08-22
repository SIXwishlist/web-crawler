package main

import (
	"golang.org/x/net/html"
	"regexp"
	"strings"
)

var domainRegex = regexp.MustCompile(`https?:\/\/([\w\d])+(\.\w+)*(:\d{0,5})*`)

type HtmlDoc interface {
	ExtractPageInfo() pageInfo
}

type htmlDoc struct {
	body    string
	domain  string
	address string
}

type HtmlDocConstructor func(string, string) HtmlDoc

type pageInfo struct {
	page   string
	links  []string
	assets []string
}

func NewHtmlDoc(body string, address string) HtmlDoc {
	domain := domainRegex.FindString(address)

	return htmlDoc{body: body, domain: domain, address: address}
}

type Condition func(*html.Node) bool

func isLink(n *html.Node) bool {
	return n.Type == html.ElementNode && n.Data == "a"
}

func isAsset(n *html.Node) bool {
	return n.Type == html.ElementNode && (n.Data == "script" || n.Data == "img")
}

func selectNodes(n *html.Node, buf []string, cond Condition, key string) []string {
	if cond(n) {
		for _, attr := range n.Attr {
			if attr.Key == key {
				buf = append(buf, attr.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		buf = selectNodes(c, buf, cond, key)
	}

	return buf
}

func filterInternalLinks(links []string, domain string) (internalLinks []string) {
	for _, link := range links {
		if strings.HasPrefix(link, domain) {
			internalLinks = append(internalLinks, link)
		} else if strings.HasPrefix(link, "/") {
			internalLinks = append(internalLinks, domain+link)
		}
	}
	return
}

func removeDuplicates(links []string) (uniqueLinks []string) {
	linksMap := make(map[string]bool)

	for _, link := range links {
		link = strings.TrimSuffix(link, "/")
		linksMap[link] = true
	}

	for link, _ := range linksMap {
		uniqueLinks = append(uniqueLinks, link)
	}

	return
}

func (this htmlDoc) ExtractPageInfo() (info pageInfo) {
	doc, err := html.Parse(strings.NewReader(this.body))
	if err != nil {
		return pageInfo{}
	}

	info.page = this.address
	info.links = this.extractInternalLinks(doc)
	info.assets = this.extractAssets(doc)

	return
}

func (this htmlDoc) extractInternalLinks(doc *html.Node) (links []string) {
	links = selectNodes(doc, links, isLink, "href")
	links = filterInternalLinks(links, this.domain)
	links = removeDuplicates(links)

	return
}

func (this htmlDoc) extractAssets(doc *html.Node) (assets []string) {
	return selectNodes(doc, assets, isAsset, "src")
}

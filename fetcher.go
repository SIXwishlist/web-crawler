package main

type Fetcher interface {
	Fetch(url string) (body HtmlDoc)
}

type fetcher struct {}

func (this fetcher) Fetch(url string) (body HtmlDoc) {
	return htmlDoc{}
}

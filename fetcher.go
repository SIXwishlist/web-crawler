package main

import (
	"fmt"
	"net/http"
)

type Fetcher interface {
	Fetch(url string) (body HtmlDoc, err error)
}

type Client interface {
	Get(url string) (*http.Response, error)
}

type fetcher struct {
	client Client
}

func (this fetcher) Fetch(url string) (doc HtmlDoc, err error) {
	response, err := this.client.Get(url)
	defer response.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("Error fetching %s: %s", url, err)
	}

	return htmlDoc{body: response.Body}, nil
}

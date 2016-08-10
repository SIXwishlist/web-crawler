package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

type Fetcher interface {
	Fetch(url string) (body string, err error)
}

type Client interface {
	Get(url string) (*http.Response, error)
}

type fetcher struct {
	client Client
}

func (this fetcher) Fetch(url string) (string, error) {
	client := this.client
	if client == nil {
		client = &http.Client{}
	}

	response, err := client.Get(url)

	if err != nil {
		return "", fmt.Errorf("Error fetching %s: %s", url, err)
	}

	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Error reading %s: %s", url, err)
	}

	if response.ContentLength > 0 {
		body = body[0:response.ContentLength]
	}
	return string(body), nil
}

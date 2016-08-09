package main

import(
	"testing"
	"net/http"
)

type testClient struct {}

type Response struct {
}

func (this Response) Read(p []byte) (n int, err error) {
	return
}

func (this Response) Close() (err error) {
	return
}

func (this testClient) Get(url string) (*http.Response, error) {
	return &http.Response{Body: Response{}}, nil
}

func TestFetch(t *testing.T) {
	testFetcher := fetcher{client: testClient{}}
	doc, _ := testFetcher.Fetch("http://tomblomfield.com")

	expectedBody := "<html></html>"
	actualBody := doc.ReadBody()
	if string(actualBody) != expectedBody {
		t.Error("Response is %s not equal %s", actualBody, expectedBody)
	}
}

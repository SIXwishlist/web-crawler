package main

import(
	"testing"
	"net/http"
	"io"
)

var (
	body = []byte("<html></html>")
)

type testClient struct {}

type ResponseBody struct {
}

func (this ResponseBody) Read(p []byte) (n int, err error) {
	copy(p, body)
	return len(p), io.EOF
}

func (this ResponseBody) Close() (err error) {
	return
}

func (this testClient) Get(url string) (*http.Response, error) {
	return &http.Response{Body: ResponseBody{}, ContentLength: int64(len(body))}, nil
}

func TestFetch(t *testing.T) {
	testFetcher := fetcher{client: testClient{}}
	body, _ := testFetcher.Fetch("http://tomblomfield.com")

	expectedBody := "<html></html>"
	if string(body) != expectedBody {
		t.Error("Response is", body, " not equal", expectedBody)
	}
}

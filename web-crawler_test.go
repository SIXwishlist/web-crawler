package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestWebCrawler(t *testing.T) {
	os.Args = []string{"", "-u=http://localhost:8080/", "-w=2"}
	output, _ = os.Create("test_output")
	defer os.Remove("test_output")

	main()

	links := []string{"http://localhost:8080/1", "http://localhost:8080/2", "http://localhost:8080/page3"}

	for _, link := range links {
		found := false

		for seenLink, _ := range seen {
			if link == seenLink {
				found = true
			}
		}

		if !found {
			t.Error(link, "has not been crawled")
		}
	}

	receivedOutput, _ := ioutil.ReadFile("test_output")

	lines := strings.Split(string(receivedOutput), "\n")
	var outputLinks []string
	foundPage := false

	for i := 0; i < len(lines); i++ {
		if lines[i] == "Page: http://localhost:8080/" {
			for j := 2; j < 5; j++ {
				outputLinks = append(outputLinks, lines[i+j])
			}
			foundPage = true
		}
	}
	if !foundPage {
		t.Error("http://localhost:8080/ has not been crawled")
	}

	var expectedLinks []string
	for i := 0; i < len(links); i++ {
		expectedLinks = append(expectedLinks, "    - "+links[i])
	}
	if !equalStringSlices(outputLinks, expectedLinks) {
		t.Error(outputLinks, "doesn't contain:", expectedLinks)
	}
}

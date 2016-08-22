package main

import(
	"testing"
	"strings"
	"os"
	"io/ioutil"
)

func TestWebCrawler(t *testing.T) {
	os.Args = []string{"", "-u=http://localhost:8080/", "-w=2"}
	output, _ = os.Create("test_output")
	defer os.Remove("test_output")

	main()

	links := []string{"http://localhost:8080/1","http://localhost:8080/2", "http://localhost:8080/page3"}
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

	expectedOutput := `Page: http://localhost:8080/
  Links:
    - http://localhost:8080/1
    - http://localhost:8080/2
    - http://localhost:8080/page3
  Assets:
    - http://www.google-analytics.com/ga.js`

	receivedOutput, _ := ioutil.ReadFile("test_output")
	if !strings.Contains(string(receivedOutput), expectedOutput) {
		t.Error("\n" + string(receivedOutput) + "\n", "doesn't contain:", "\n" + expectedOutput)
	}
}

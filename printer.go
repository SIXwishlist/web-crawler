package main

import (
	"io"
)

type pageResult struct {
	page string
	links []string
}

func Printer(output io.Writer, results <-chan pageResult) {
	for result := range results {
		io.WriteString(output, "Page: " + result.page + "\n")
		io.WriteString(output, "  Links:\n")

		for _, link := range result.links {
			io.WriteString(output, "    - " + link + "\n")
		}
	}
}

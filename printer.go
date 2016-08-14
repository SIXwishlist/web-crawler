package main

import (
	"io"
)

func Printer(output io.Writer, results <-chan pageInfo) {
	for result := range results {
		io.WriteString(output, "Page: " + result.page + "\n")
		io.WriteString(output, "  Links:\n")

		for _, link := range result.links {
			io.WriteString(output, "    - " + link + "\n")
		}
	}
}

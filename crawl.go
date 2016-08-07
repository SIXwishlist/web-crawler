package main

import(
	"fmt"
)

func crawl(jobs <-chan string, done <-chan bool, foundLinks chan<- string) {
	for {
		select {
		case job := <-jobs:
			foundLinks <- "foo"
			foundLinks <- "bar"
			foundLinks <- "baz"
		case flag := <-done:
			fmt.Println("Terminating")
			return
		}
	}
}

package main

type Worker struct {
	fetcher    Fetcher
	newHtmlDoc HtmlDocConstructor
}

type WorkerConstructor func() *Worker

func NewWorker() *Worker {
	return &Worker{fetcher: fetcher{}, newHtmlDoc: NewHtmlDoc}
}

func (this Worker) Start(id int, unseenLinks <-chan string, foundLinks chan<- []string, results chan<- pageResult) {
	for link := range unseenLinks {
		links := this.extractLinks(link)
		result := pageResult{page: link, links: links}

		go func() { foundLinks <- links }()
		go func() { results <- result }()
	}
}

func (this Worker) extractLinks(link string) (links []string) {
	body, err := this.fetcher.Fetch(link)
	if err != nil {
		return
	}
	doc := this.newHtmlDoc(body, link)
	links = doc.ExtractInternalLinks()
	return
}

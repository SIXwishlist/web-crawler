package main

type Worker struct {
	fetcher    Fetcher
	newHtmlDoc HtmlDocConstructor
}

type WorkerConstructor func() *Worker

func NewWorker() *Worker {
	return &Worker{fetcher: fetcher{}, newHtmlDoc: NewHtmlDoc}
}

func (this Worker) Start(id int, unseenLinks <-chan string, foundLinks chan<- []string, results chan<- pageInfo) {
	for link := range unseenLinks {
		info := this.extractPageInfo(link)

		go func() { foundLinks <- info.links }()
		go func() { results <- info }()
	}
}

func (this Worker) extractPageInfo(link string) (info pageInfo) {
	body, err := this.fetcher.Fetch(link)
	if err != nil {
		return
	}
	doc := this.newHtmlDoc(body, link)
	info = doc.ExtractPageInfo()
	return
}

# Concurrent Web Crawler in go

Crawls a domain printing out the site map with static assets for each page and links between pages. Ignores the links to the "outside".

## How to use

Clone this repo and run `go install`. Run `$GOPATH/web-crawler -u <URL> -w <NO. OF WORKERS>`

## How to run tests

Clone this repo and run `go test`.

## Requirements

- Number of workers is specified in a command line flag

```
web-crawler -u http://tomblomfield.com -w 20
```

- Only crawls the links within the given domain
- Outputs a site map with static assets for each page
- Outputs links between pages

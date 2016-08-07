# Concurrent Web Crawler in go

Crawls a domain printing out the site map with static assets for each page and links between pages. Ignores the links to the "outside".

## How to use

TODO

## How to run tests

TODO

## Requirements

- Number of workers is specified in a command line flag

```
web-crawler -u google.com -w 20
```

- Only crawls the links within the given domain
- Outputs a site map with static assets for each page
- Outputs links between pages

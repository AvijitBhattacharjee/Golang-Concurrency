package pkg

import (
	"fmt"
	"sync"
)

type Crawler struct {
	visited map[string]bool
	mu      sync.Mutex
}

func (c *Crawler) crawl(url string) {

	c.mu.Lock()

	if c.visited[url] {

		c.mu.Unlock()
		return
	}

	c.visited[url] = true

	c.mu.Unlock()

	fmt.Println("crawling:", url)

	links := fakeFetch(url)

	for _, link := range links {
		go c.crawl(link)
	}
}

func fakeFetch(url string) []string {

	graph := map[string][]string{
		"google.com": {"gmail.com", "maps.com"},
		"gmail.com":  {"docs.com"},
		"maps.com":   {"docs.com"},
		"docs.com":   {},
	}

	return graph[url]
}

func WebCrawler() {

	crawler := &Crawler{
		visited: make(map[string]bool),
	}

	crawler.crawl("google.com")

	fmt.Scanln()
}

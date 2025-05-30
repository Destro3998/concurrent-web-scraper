package main

import (
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org",
		"https://example.com",
		"https://github.com",
		"https://news.ycombinator.com",
	}

	var wg sync.WaitGroup
	jobs := make(chan string)

	// Start 5 workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				scrape(url)
			}
		}()
	}

	// Send URLs into the channel
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
}

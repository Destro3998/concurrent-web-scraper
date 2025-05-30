package main

import (
	"os"
	"sync"
)

func main() {
	// create and initialize HTML file
	os.WriteFile("output.html", []byte("<html><body><h1>Scraped Links</h1>\n"), 0644)

	// URL to scrape
	urls := []string{
		"https://golang.org",
		"https://github.com",
		"https://news.ycombinator.com",
	}

	var wg sync.WaitGroup
	jobs := make(chan string)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range jobs {
				scrape(url)
			}
		}()
	}

	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()

	// close HTML file
	f, _ := os.OpenFile("output.html", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString("</body></html>")
}

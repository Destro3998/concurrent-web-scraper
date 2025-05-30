package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var mu sync.Mutex // writing to file concurrently

func scrape(url string) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		log.Println("Error fetching:", url, err)
		return
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("Error reading body:", err)
		return
	}

	// To open HTML file for appending
	mu.Lock()
	file, err := os.OpenFile("output.html", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error opening output file:", err)
		mu.Unlock()
		return
	}

	// display links in HTML format
	fmt.Fprintf(file, "<h2>Links from %s</h2>\n<ul>\n", url)
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Fprintf(file, "<li><a href='%s'>%s</a></li>\n", href, href)
		}
	})
	fmt.Fprintln(file, "</ul><hr>")
	file.Close()
	mu.Unlock()
}

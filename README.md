# Concurrent Web Scraper using Go

This is a simple and efficient concurrent web scraping tool built using Go. It shows how to use goroutines, channels, and worker pools to fetch and extract links (`<a href="">`) from multiple web pages in parallel.


## Features

- Fetches web pages concurrently using a worker pool
- Extracts and prints all links (`href`) from each page
- Uses `goquery` for HTML parsing (jQuery-like)
- Clean modular with concurrency patterns


## How to Run

1. Clone the repository
- git clone https://github.com/Destro3998/concurrent-web-scraper.git
- cd concurrent-web-scraper

2. Make sure Go is installed
- go version

3. Run the scraper
- go run main.go scraper.go
- open output.html  


## Performance: Sequential vs Concurrent

Scraping multiple websites concurrently significantly reduces the total execution time compared to a sequential approach. A sequential scraper took 10.4 seconds to complete the same task that a concurrent scraper handled much faster.

### 🚀 Concurrent Scraper (3 Goroutines)

```bash
Scraping: https://golang.org
Scraping: https://github.com
Scraping: https://news.ycombinator.com

Elapsed time: 2.3s - nearly 4-5x faster.
# Concurrent Web Scraper in Go 🕸️⚙️

This is a simple and efficient concurrent web scraping tool built using Go. It demonstrates how to use goroutines, channels, and worker pools to fetch and extract links (`<a href="">`) from multiple web pages in parallel.

## 🔧 Features

- Fetches web pages concurrently using a worker pool
- Extracts and prints all links (`href`) from each page
- Uses `goquery` for HTML parsing (jQuery-like)
- Clean modular code with concurrency patterns

## 📁 Project Structure

o-concurrent-scraper/
├── main.go # Entry point, handles concurrency and workers
├── scraper.go # Contains the HTML scraping logic
├── go.mod # Go module file
├── go.sum # Dependency verification file


## 🚀 How to Run

### 1. Clone the repo

```bash
git clone https://github.com/Destro3998/concurrent-web-scraper.git
cd concurrent-web-scraper

### 2. Make sure Go is Installed

go version

### 3. Run the scraper

go run main.go scraper.go

✏️ Sample Output

Links found on https://golang.org
 → https://pkg.go.dev
 → https://play.golang.org
 → ...

Links found on https://github.com
 → https://github.com/login
 → https://github.com/signup
 → ...



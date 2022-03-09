package main

import (
	"log"
)

func main() {
	// Initialize crawler
	crawler := NewCrawler(4)

	log.Printf("Crawler: %+v\n", crawler)
}

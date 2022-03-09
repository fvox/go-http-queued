package main

import (
	"log"
)

func main() {
	// Initialize crawler
	crawler := NewCrawler(4)
	crawler.SpawnWorkers()

	crawler.Enqueue("http://github.com")

	log.Printf("Crawler: %+v\n", crawler)

	crawler.Wait()

	log.Println("Finished")
}

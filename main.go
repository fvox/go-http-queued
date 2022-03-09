package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize crawler
	crawler := NewCrawler(4)
	crawler.SpawnWorkers()

	// Enqueue jobs
	req, _ := http.NewRequest("GET", "http://github.com/", nil)

	crawler.Enqueue(req, func(res *http.Response) {
		log.Printf("Github response: %+v\n", res)
	})

	log.Printf("Crawler: %+v\n", crawler)

	crawler.Wait()

	log.Println("Finished")
}

package main

type Crawler struct {
	MaxWorkers int
}

func NewCrawler(maxWorkers int) *Crawler {
	return &Crawler{MaxWorkers: maxWorkers}
}

package main

type Crawler struct {
	MaxWorkers int
	jobChan    chan Job
}

func NewCrawler(maxWorkers int) *Crawler {
	jobChan := make(chan Job)

	return &Crawler{
		MaxWorkers: maxWorkers,
		jobChan:    jobChan,
	}
}

func (c *Crawler) Enqueue(url string) {
	go func() {
		job := Job{Url: url}
		c.jobChan <- job
	}()
}

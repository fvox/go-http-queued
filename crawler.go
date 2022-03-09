package main

import (
	"log"
	"sync"
)

type Crawler struct {
	MaxWorkers int
	jobChan    chan Job
	waitGroup  *sync.WaitGroup
}

func NewCrawler(maxWorkers int) *Crawler {
	jobChan := make(chan Job)
	wg := &sync.WaitGroup{}

	return &Crawler{
		MaxWorkers: maxWorkers,
		jobChan:    jobChan,
		waitGroup:  wg,
	}
}

func (c *Crawler) SpawnWorkers() {
	for i := 1; i <= c.MaxWorkers; i++ {
		go func(workerId int) {
			for job := range c.jobChan {
				log.Printf("Job: %+v\n", job)

				c.waitGroup.Done()
			}
		}(i)
	}
}

func (c *Crawler) Enqueue(url string) {
	c.waitGroup.Add(1)

	go func() {
		job := Job{Url: url}
		c.jobChan <- job
	}()
}

func (c *Crawler) Wait() {
	c.waitGroup.Wait()
	close(c.jobChan)
}

package main

import (
	"net/http"
	"sync"
)

type Crawler struct {
	MaxWorkers int
	jobChan    chan Job
	waitGroup  *sync.WaitGroup
	Client     *http.Client
}

func NewCrawler(maxWorkers int) *Crawler {
	jobChan := make(chan Job)
	wg := &sync.WaitGroup{}
	client := &http.Client{}

	return &Crawler{
		MaxWorkers: maxWorkers,
		Client:     client,
		jobChan:    jobChan,
		waitGroup:  wg,
	}
}

func (c *Crawler) SpawnWorkers() {
	for i := 1; i <= c.MaxWorkers; i++ {
		go func(workerId int) {
			for job := range c.jobChan {
				job.WorkerId = workerId
				job.Perform()
				c.waitGroup.Done()
			}
		}(i)
	}
}

func (c *Crawler) Enqueue(req *http.Request, action func(res *http.Response)) {
	c.waitGroup.Add(1)
	go func() {
		job := Job{Crawler: c, Req: req, Action: action}
		c.jobChan <- job
	}()
}

func (c *Crawler) Wait() {
	c.waitGroup.Wait()
	close(c.jobChan)
}

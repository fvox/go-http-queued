package main

import "net/http"

type Job struct {
	*Crawler
	WorkerId int
	Req      *http.Request
	Action   func(res *http.Response)
}

func (j *Job) Perform() {
	// Do HTTP request
	res, _ := j.Crawler.Client.Do(j.Req)
	defer res.Body.Close()

	// Perform job action
	j.Action(res)
}

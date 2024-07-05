package model

// url: https://www.instagram.com/kyliejenner
// url_timeout: 1
// runtime: 1
// parallel_rqquests: 1

type JobRequest struct {
	Url              string `json:"url"`
	UrlTimeout       int    `json:"url_timeout"`
	Runtime          int    `json:"runtime"`
	ParallelRequests int    `json:"parallel_requests"`
}

package model

import (
	"time"

	"github.com/google/uuid"
)

type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
)

type Job struct {
	ID               uuid.UUID
	Status           Status
	CreatedAt        time.Time
	StartedAt        time.Time
	CompletedAt      time.Time
	Uri              string
	Timeout          int
	Runtime          int
	ParallelRequests int
	Result           *Result
	Debug            bool
}

func NewJob(uri string, timeout, runtime, parallelRequests int, status Status) *Job {
	return &Job{
		ID:               uuid.New(),
		Status:           status,
		CreatedAt:        time.Now(),
		Uri:              uri,
		Timeout:          timeout,
		Runtime:          runtime,
		ParallelRequests: parallelRequests,
		Result:           NewResult(),
		Debug:            true,
	}

}

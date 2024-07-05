package model

import (
	"sync"

	"github.com/google/uuid"
)

type Jobs struct {
	mu   sync.RWMutex
	jobs []*Job
}

func NewJobs() *Jobs {
	return &Jobs{
		jobs: make([]*Job, 0),
	}
}

// Add adds a todo to the list
func (j *Jobs) Add(uri string, timeout, runtime, parallelRequests int, status Status) *Job {
	j.mu.Lock()
	defer j.mu.Unlock()

	job := NewJob(uri, timeout, runtime, parallelRequests, status)
	j.jobs = append(j.jobs, job)
	return job
}

// Remove removes a todo from the list
func (j *Jobs) Remove(id uuid.UUID) {
	j.mu.Lock()
	defer j.mu.Unlock()

	index := j.indexOf(id)
	if index == -1 {
		return
	}
	j.jobs = append((j.jobs)[:index], (j.jobs)[index+1:]...)
}

// All returns a copy of the list of todos
func (j *Jobs) All() []*Job {
	j.mu.RLock()
	defer j.mu.RUnlock()

	list := make([]*Job, len(j.jobs))
	copy(list, j.jobs)
	return list
}

// Get returns a todo by id
func (j *Jobs) Get(id uuid.UUID) *Job {
	j.mu.RLock()
	defer j.mu.RUnlock()

	index := j.indexOf(id)
	if index == -1 {
		return nil
	}
	return (j.jobs)[index]
}

// indexOf returns the index of the todo with the given id or -1 if not found
func (j *Jobs) indexOf(id uuid.UUID) int {
	for i, job := range j.jobs {
		if job.ID == id {
			return i
		}
	}
	return -1
}

func (j *Jobs) Next() *Job {
	j.mu.Lock()
	defer j.mu.Unlock()

	if len(j.jobs) == 0 {
		return nil
	}
	for _, job := range j.jobs {
		if job.Status == Pending {
			job.Status = Running
			return job
		}
	}
	return nil
}

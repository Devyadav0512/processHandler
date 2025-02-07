package queue

import (
	"processhandler/models"
)

type Queue struct {
	queue   chan models.Job
	workers int
}

func NewQueue(workersCount int) *Queue {
	return &Queue{
		queue:   make(chan models.Job),
		workers: workersCount,
	}
}

func (q *Queue) Enqueue(msg models.Job) {
	q.queue <- msg
}

func (q *Queue) Dequeue() models.Job {
	return <-q.queue
}

func (q *Queue) Subscribe(ExecutionHandler func(*models.Job) bool, expQ *Queue) {
	for range q.workers {
		go q.Worker(ExecutionHandler, expQ)
	}
}

func (q *Queue) Worker(ExecutionHandler func(*models.Job) bool, expQ *Queue) {
	for job := range q.queue {
		res := ExecutionHandler(&job)
		if !res {
			expQ.Enqueue(job)
		}
	}
}

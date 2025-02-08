package queue

import (
	"fmt"
	"processhandler/models"
	"sync"
)

type Queue struct {
	qType          string
	queue          chan models.Job
	workers        int
	workersWaitGrp sync.WaitGroup
}

func NewQueue(workersCount int, queueType string) *Queue {
	return &Queue{
		qType:          queueType,
		queue:          make(chan models.Job),
		workers:        workersCount,
		workersWaitGrp: sync.WaitGroup{},
	}
}

func (q *Queue) Enqueue(msg models.Job) {
	q.queue <- msg
}

func (q *Queue) Close() {
	close(q.queue)
}

func (q *Queue) Subscribe(ExecutionHandler func(models.Job) bool, expQ *Queue, wg *sync.WaitGroup) {
	defer func() {
		q.workersWaitGrp.Wait()
		fmt.Println(q.qType, " queue, closed")
		wg.Done()
	}()
	for range q.workers {
		q.workersWaitGrp.Add(1)
		go q.Worker(ExecutionHandler, expQ)
	}
}

func (q *Queue) Worker(ExecutionHandler func(models.Job) bool, expQ *Queue) {
	defer func() {
		q.workersWaitGrp.Done()
	}()
	for job := range q.queue {
		res := ExecutionHandler(job)
		if !res {
			expQ.Enqueue(job)
		}
	}
}

package jobsHandler

import (
	"processhandler/jobHandler"
	"processhandler/models"
	"processhandler/queue"
	"sync"
)

var jobWg sync.WaitGroup
var excpWg sync.WaitGroup

func InitAllJobHandler(jobs []models.Job) {
	FastJobQueue := queue.NewQueue(5, "fast")
	MedJobQueue := queue.NewQueue(5, "med")
	SlowJobQueue := queue.NewQueue(5, "slow")
	ExceptionQueue := queue.NewQueue(1, "excp")

	defer func() {
		jobWg.Wait()
		ExceptionQueue.Close()
		excpWg.Wait()
	}()

	jobWg.Add(6)
	excpWg.Add(1)

	go FastJobQueue.Subscribe(jobHandler.FastJobHandler, ExceptionQueue, &jobWg)
	go MedJobQueue.Subscribe(jobHandler.MedJobHandler, ExceptionQueue, &jobWg)
	go SlowJobQueue.Subscribe(jobHandler.SlowJobHandler, ExceptionQueue, &jobWg)
	go ExceptionQueue.Subscribe(jobHandler.ExcpJobHandler, ExceptionQueue, &excpWg)

	go AddJobs(jobs, FastJobQueue, "fast")
	go AddJobs(jobs, MedJobQueue, "medium")
	go AddJobs(jobs, SlowJobQueue, "slow")

}

func AddJobs(jobs []models.Job, queue *queue.Queue, jobType string) {
	defer func() {
		queue.Close()
		jobWg.Done()
	}()

	for _, job := range jobs {
		if job.Type == jobType {
			queue.Enqueue(job)
		}
	}
}

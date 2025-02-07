package main

import (
	"fmt"
	"processhandler/jobHandler"
	"processhandler/models"
	"processhandler/queue"
	"sync"
)

var addJobwg sync.WaitGroup

func InitAllJobHandler(jobs []models.Job) {
	FastJobQueue := queue.NewQueue(5)
	MedJobQueue := queue.NewQueue(5)
	SlowJobQueue := queue.NewQueue(5)
	ExceptionQueue := queue.NewQueue(1)

	defer addJobwg.Wait()
	addJobwg.Add(3)

	go AddJobs(jobs, FastJobQueue, "fast")
	go AddJobs(jobs, MedJobQueue, "medium")
	go AddJobs(jobs, SlowJobQueue, "slow")

	go FastJobQueue.Subscribe(jobHandler.FastJobHandler, ExceptionQueue)
	go MedJobQueue.Subscribe(jobHandler.MedJobHandler, ExceptionQueue)
	go SlowJobQueue.Subscribe(jobHandler.SlowJobHandler, ExceptionQueue)
	go ExceptionQueue.Subscribe(jobHandler.ExcpJobHandler, ExceptionQueue)
}

func AddJobs(jobs []models.Job, queue *queue.Queue, jobType string) {
	defer func() {
		fmt.Println(jobType + " job addition success")
		addJobwg.Done()
	}()

	for _, job := range jobs {
		if job.Type == jobType {
			queue.Enqueue(job)
		}
	}
}

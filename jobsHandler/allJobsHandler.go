package jobsHandler

import (
	"fmt"
	"processhandler/jobHandler"
	"processhandler/models"
	"processhandler/queue"
	"sync"
)

var jobWg sync.WaitGroup
var excpWg sync.WaitGroup

func InitAllJobHandler(jobs []models.Job) []models.JobResult {
	jobsResponse := models.Response{}

	FastJobQueue := queue.NewQueue(5, "fast")
	MedJobQueue := queue.NewQueue(5, "med")
	SlowJobQueue := queue.NewQueue(5, "slow")
	ExceptionQueue := queue.NewQueue(1, "excp")

	queueArr := []queues{
		{FastJobQueue, &jobWg, "fast", jobHandler.FastJobHandler, ExceptionQueue, true},
		{MedJobQueue, &jobWg, "medium", jobHandler.MedJobHandler, ExceptionQueue, true},
		{SlowJobQueue, &jobWg, "slow", jobHandler.SlowJobHandler, ExceptionQueue, true},
		{ExceptionQueue, &excpWg, "excp", jobHandler.ExcpJobHandler, ExceptionQueue, false},
	}

	defer func() {
		jobWg.Wait()
		ExceptionQueue.Close()
		CalculateStats(queueArr, &jobsResponse)
		excpWg.Wait()
	}()

	jobWg.Add(6)
	excpWg.Add(1)

	for _, q := range queueArr {
		go q.queue.Subscribe(q.jobHandler, q.excpQueue, q.waitGrp)
		if q.isJobQueue {
			go AddJobs(jobs, q.queue, q.jobType)
		}
	}

	return jobsResponse.JobsResult
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

func CalculateStats(queueArr []queues, jobResponse *models.Response) {
	for _, q := range queueArr {
		if q.isJobQueue {
			completedJobs, failedJobs := q.queue.Stats()
			jobResponse.CompletedJobs += completedJobs
			jobResponse.FailedJobs += failedJobs
		}
	}
	fmt.Println("Total completed jobs: ", jobResponse.CompletedJobs, " failed jobs: ", jobResponse.FailedJobs)
}

package jobhandler

import (
	"processhandler/models"
	"sync"
)

// fast job 1 -> push
// fast job 1 execution
// about to complete -> remove event from queue
// same time -> job handler tries to add event to the queue

var fastJobQueue []models.Job
var mediumJobsQueue []models.Job
var slowJobsQueue []models.Job

var fastJobLock sync.Mutex
var MedJobLock sync.Mutex
var SlowJobLock sync.Mutex

var processes chan string = make(chan string, 5)
var wg sync.WaitGroup

// job addition -> 'completed'
// slow and med -> go routine
// fast -> go routine

func Init(jobs []models.Job) {
	for i := range jobs {
		if jobs[i].Type == "fast" {
			fastJobQueue = append(fastJobQueue, jobs[i])
		} else if jobs[i].Type == "medium" {
			mediumJobsQueue = append(mediumJobsQueue, jobs[i])
		} else if jobs[i].Type == "slow" {
			slowJobsQueue = append(slowJobsQueue, jobs[i])
		}
	}

	wg.Wait()
	// Close channel when job processing done;
}

func AddFastJob(fastJob models.Job) {
	fastJobLock.Lock()
	defer fastJobLock.Unlock()
	fastJobQueue = append(fastJobQueue, fastJob)
	if len(fastJobQueue) == 1 && len(processes) < cap(processes) {
		processes <- "fastJob" + fastJob.Id
		wg.Add(1)
		go FastJobHandler()
	}
}

func AddMedJob(medJob models.Job) {
	MedJobLock.Lock()
	defer MedJobLock.Unlock()
	mediumJobsQueue = append(mediumJobsQueue, medJob)
	if len(mediumJobsQueue) == 1 && len(processes) < cap(processes) {
		processes <- "medJob" + medJob.Id
		wg.Add(1)
		go MedJobHandler()
	}
}

func AddSlowJob(slowJob models.Job) {
	SlowJobLock.Lock()
	defer SlowJobLock.Unlock()
	slowJobsQueue = append(slowJobsQueue, slowJob)
	if len(slowJobsQueue) == 1 && len(processes) < cap(processes) {
		processes <- "slowJob" + slowJob.Id
		wg.Add(1)
		go SlowJobHandler()
	}
}

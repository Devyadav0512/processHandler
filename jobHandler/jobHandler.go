package jobhandler

import (
	"fmt"
	"processhandler/models"
	"sync"
)

var fastJobQueue []models.Job
var mediumJobsQueue []models.Job
var slowJobsQueue []models.Job

var fastJobLock sync.Mutex
var MedJobLock sync.Mutex
var SlowJobLock sync.Mutex

var fastProcesses chan string = make(chan string, 5)
var medProcesses chan string = make(chan string, 5)
var slowProcesses chan string = make(chan string, 5)
var fwg sync.WaitGroup
var mwg sync.WaitGroup
var swg sync.WaitGroup

var addJobwg sync.WaitGroup
var jobAddChan chan string = make(chan string)
var fastJobAdd bool = false
var medJobAdd bool = false
var slowJobAdd bool = false

func Init(jobs []models.Job) {
	go AddFastJobs(&jobs)
	go AddMedJobs(&jobs)
	go AddSlowJobs(&jobs)

	addJobwg.Add(3)

	go JobAdditionChecker()

	addJobwg.Wait()

	go func() {
		fwg.Wait()
		close(fastProcesses)
	}()

	go func() {
		mwg.Wait()
		close(medProcesses)
	}()

	go func() {
		swg.Wait()
		close(slowProcesses)
	}()
}

func JobAdditionChecker() {
	for msg := range jobAddChan {
		if msg == "fastJob" {
			fastJobAdd = true
			if medJobAdd && slowJobAdd {
				close(jobAddChan)
			}
		} else if msg == "medJob" {
			medJobAdd = true
			if fastJobAdd && slowJobAdd {
				close(jobAddChan)
			}
		} else if msg == "slowJob" {
			slowJobAdd = true
			if medJobAdd && fastJobAdd {
				close(jobAddChan)
			}
		}
	}
}

func AddFastJobs(jobs *[]models.Job) {
	defer addJobwg.Done()
	defer func() {
		fmt.Println("Fast Job Addition Success")
		jobAddChan <- "fastJob"
	}()

	for _, job := range *jobs {
		if job.Type == "fast" {
			fastJobLock.Lock()
			defer fastJobLock.Unlock()
			fastJobQueue = append([]models.Job{job}, fastJobQueue...)
			if len(fastJobQueue) == 1 {
				if len(fastProcesses) < cap(fastProcesses) {
					fastProcesses <- "fastJob" + job.Id
					fwg.Add(1)
					go FastJobHandler()
				} else if medJobAdd && len(medProcesses) < cap(medProcesses) {
					medProcesses <- "fastJob" + job.Id
					mwg.Add(1)
					go FastJobHandler()
				} else if slowJobAdd && len(slowProcesses) < cap(slowProcesses) {
					slowProcesses <- "fastJob" + job.Id
					swg.Add(1)
					go FastJobHandler()
				}
			}
		}
	}
}

func AddMedJobs(jobs *[]models.Job) {
	defer addJobwg.Done()
	defer func() {
		fmt.Println("Medium Job Addition Success")
		jobAddChan <- "medJob"
	}()

	for _, job := range *jobs {
		if job.Type == "medium" {
			MedJobLock.Lock()
			defer MedJobLock.Unlock()
			mediumJobsQueue = append([]models.Job{job}, mediumJobsQueue...)
			if len(mediumJobsQueue) == 1 {
				if len(medProcesses) < cap(medProcesses) {
					medProcesses <- "medJob" + job.Id
					mwg.Add(1)
					go MedJobHandler()
				} else if fastJobAdd && len(fastProcesses) < cap(fastProcesses) {
					fastProcesses <- "medJob" + job.Id
					fwg.Add(1)
					go MedJobHandler()
				} else if slowJobAdd && len(slowProcesses) < cap(slowProcesses) {
					slowProcesses <- "medJob" + job.Id
					swg.Add(1)
					go MedJobHandler()
				}
			}
		}
	}
}

func AddSlowJobs(jobs *[]models.Job) {
	defer addJobwg.Done()
	defer func() {
		fmt.Println("Slow Job Addition Success")
		jobAddChan <- "slowJob"
	}()

	for _, job := range *jobs {
		if job.Type == "slow" {
			SlowJobLock.Lock()
			defer SlowJobLock.Unlock()
			slowJobsQueue = append([]models.Job{job}, slowJobsQueue...)
			if len(slowJobsQueue) == 1 {
				if len(slowProcesses) < cap(slowProcesses) {
					slowProcesses <- "slowJob" + job.Id
					swg.Add(1)
					go SlowJobHandler()
				} else if fastJobAdd && len(fastProcesses) < cap(fastProcesses) {
					fastProcesses <- "slowJob" + job.Id
					fwg.Add(1)
					go SlowJobHandler()
				} else if medJobAdd && len(medProcesses) < cap(medProcesses) {
					medProcesses <- "slowJob" + job.Id
					mwg.Add(1)
					go SlowJobHandler()
				}
			}
		}
	}
}

package jobsHandler

import (
	"processhandler/models"
	"processhandler/queue"
	"sync"
)

type queues struct {
	queue      *queue.Queue
	waitGrp    *sync.WaitGroup
	jobType    string
	jobHandler func(models.Job) bool
	excpQueue  *queue.Queue
	isJobQueue bool
}

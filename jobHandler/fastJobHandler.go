package jobHandler

import (
	"fmt"
	"processhandler/models"
	"time"
)

func FastJobHandler(job *models.Job) bool {
	fmt.Println("Job "+job.Id+" started - Type: "+job.Type, job.Payload)
	time.Sleep(1 * time.Second)
	fmt.Println("Job " + job.Id + " completed - Duration: 1000ms")
	return true
}

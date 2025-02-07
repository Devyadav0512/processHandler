package jobHandler

import (
	"fmt"
	"processhandler/models"
	"time"
)

func SlowJobHandler(job *models.Job) bool {
	fmt.Println("Job "+job.Id+" started - Type: "+job.Type, job.Payload)
	time.Sleep(5 * time.Second)
	fmt.Println("Job " + job.Id + " completed - Duration: 5000ms")
	return true
}

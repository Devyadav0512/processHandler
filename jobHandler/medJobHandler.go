package jobHandler

import (
	"fmt"
	"processhandler/models"
	"time"
)

func MedJobHandler(job *models.Job) bool {
	fmt.Println("Job "+job.Id+" started - Type: "+job.Type, job.Payload)
	time.Sleep(3 * time.Second)
	fmt.Println("Job " + job.Id + " completed - Duration: 3000ms")
	return true
}

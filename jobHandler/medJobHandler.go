package jobHandler

import (
	"fmt"
	"processhandler/models"
	"time"
)

func MedJobHandler(job models.Job) bool {
	fmt.Println("Job " + job.Id + " started - Type: " + job.Type)
	start := time.Now()
	time.Sleep(3 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Job "+job.Id+" completed - Duration: ", elapsed)
	return true
}

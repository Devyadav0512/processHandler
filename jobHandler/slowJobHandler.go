package jobHandler

import (
	"fmt"
	"processhandler/models"
	"time"
)

func SlowJobHandler(job models.Job) bool {
	fmt.Println("Job " + job.Id + " started - Type: " + job.Type)
	start := time.Now()
	time.Sleep(5 * time.Second)
	elapsed := time.Since(start)
	fmt.Println("Job "+job.Id+" completed - Duration: ", elapsed)
	return true
}

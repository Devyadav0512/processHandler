package jobHandler

import (
	"fmt"
	"processhandler/models"
)

func ExcpJobHandler(job models.Job) bool {
	fmt.Println("Job " + job.Id + " failed ---> " + job.Type)
	return true
}

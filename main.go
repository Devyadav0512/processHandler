package main

import (
	"fmt"
	"processhandler/jobsHandler"
	"processhandler/models"
	"time"
)

func main() {
	jobs := []models.Job{
		{
			Id:        "1",
			Type:      "fast",
			Payload:   models.JobInput{Data: "reverse this string"},
			CreatedAt: time.Now().Unix(),
			Status:    "pending",
		},
		{
			Id:        "2",
			Type:      "medium",
			Payload:   models.JobInput{Data: "[1,5,2,7,3]"},
			CreatedAt: time.Now().Unix(),
			Status:    "pending",
		},
		{
			Id:        "3",
			Type:      "slow",
			Payload:   models.JobInput{Data: "50"}, // ex : Calculate prime numbers up to 50
			CreatedAt: time.Now().Unix(),
			Status:    "pending",
		},
		{
			Id:        "4",
			Type:      "fast",
			Payload:   models.JobInput{Data: "slice this string"},
			CreatedAt: time.Now().Unix(),
			Status:    "pending",
		},
	}

	jobsStatus := jobsHandler.InitAllJobHandler(jobs)
	fmt.Println("Jobs Status: ", jobsStatus)
}

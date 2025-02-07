package main

import (
	"fmt"
	jobhandler "processhandler/jobHandler"
	"processhandler/models"
)

func main() {
	jobs := []models.Job{
		{
			Id:      "1",
			Type:    "fast",
			Payload: models.JobInput{Data: "reverse this string"},
		},
		{
			Id:      "2",
			Type:    "medium",
			Payload: models.JobInput{Data: "[1,5,2,7,3]"},
		},
		{
			Id:      "3",
			Type:    "slow",
			Payload: models.JobInput{Data: "50"}, // ex : Calculate prime numbers up to 50
		},
		{
			Id:      "4",
			Type:    "fast",
			Payload: models.JobInput{Data: "slice this string"},
		},
	}

	jobhandler.Init(jobs)
	fmt.Println("Hello world")
}

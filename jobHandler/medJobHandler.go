package jobhandler

import "fmt"

func MedJobHandler() {
	fmt.Println("Hello med job")
	mediumJobsQueue = mediumJobsQueue[:len(mediumJobsQueue)-1]
}

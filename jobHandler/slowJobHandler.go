package jobhandler

import "fmt"

func SlowJobHandler() {
	fmt.Println("Hello slow job")
	slowJobsQueue = slowJobsQueue[:len(slowJobsQueue)-1]
}

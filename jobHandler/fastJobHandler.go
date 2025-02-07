package jobhandler

import "fmt"

func FastJobHandler() {
	fmt.Println("Hello fast job")
	fastJobQueue = fastJobQueue[:len(fastJobQueue)-1]
}

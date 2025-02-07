package jobhandler

func FastJobHandler() {
	// process job
	fastJobLock.Lock()
	defer fastJobLock.Unlock()
	if len(fastJobQueue) > 0 && len(processes) < cap(processes) {
		fastJob := fastJobQueue[len(fastJobQueue)-1]
		processes <- "fastJob" + fastJob.Id
		wg.Add(1)
		go FastJobHandler()
	}
}

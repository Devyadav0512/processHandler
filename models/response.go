package models

type Response struct {
	CompletedJobs int
	FailedJobs    int
	JobsResult    []JobResult
}

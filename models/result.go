package models

type JobResult struct {
	JobId     string
	StartTime int    // Processing start time
	EndTime   int    // Processing end time
	Duration  int    // Processing duration in ms
	Status    string // 'completed'|'failed'
	Error     string // Error message if failed
}

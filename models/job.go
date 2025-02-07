package models

type Job struct {
	Id        string // Unique identifier
	Type      string // 'fast'|'medium'|'slow';
	Payload   JobInput
	CreatedAt int    // Unix timestamp
	Status    string // 'pending'|'processing'|'completed'|'failed'

}

type JobInput struct {
	Data string // Job input data
}

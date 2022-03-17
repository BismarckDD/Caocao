package c2enum

// Priority of the event handler
type Priority uint8

// Priorities
const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)

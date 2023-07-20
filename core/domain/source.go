package domain

import "go.uber.org/zap"

// EventSQS represents a process.
type EventSQS struct {
	ID            string
	Retry         string
	Records       Events
	OriginalEvent interface{}
	Log           *zap.SugaredLogger
}

// Source represents a source of events.
type Source interface {
	Consume() <-chan *EventSQS
	Processed(e *EventSQS) error
	Close() error
}

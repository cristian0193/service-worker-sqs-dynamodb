package events

import (
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/dataproviders/awsdynamodb/repository/events"
)

type IEventCaseUses interface {
	GetID(ID string) (*domain.Events, error)
}

// EventCaseUses encapsulates all the data necessary for the implementation of the EventsRepository.
type EventCaseUses struct {
	eventRepository events.IEventRepository
}

// NewEventUseCases instance the repository usecases.
func NewEventUseCases(er events.IEventRepository) *EventCaseUses {
	return &EventCaseUses{
		eventRepository: er,
	}
}

// GetID return the event by ID.
func (es *EventCaseUses) GetID(ID string) (*domain.Events, error) {
	return es.eventRepository.GetID(ID)
}

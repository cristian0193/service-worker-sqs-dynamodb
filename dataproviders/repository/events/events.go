package events

import (
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/core/domain/exceptions"
	"service-worker-sqs-dynamodb/dataproviders/awsdynamodb"
)

type IEventsRepository interface {
	GetID(ID string) (*domain.Events, error)
	Insert(events *domain.Events) error
}

// EventsRepository encapsulates all the data needed to the persistence in the event table.
type EventsRepository struct {
	db *awsdynamodb.ClientDynamoDB
}

// NewEventsRepository instance the connection to the dynamodb.
func NewEventsRepository(db *awsdynamodb.ClientDynamoDB) *EventsRepository {
	return &EventsRepository{
		db: db,
	}
}

// GetID return the event by ID.
func (er *EventsRepository) GetID(ID string) (*domain.Events, error) {
	event, err := er.db.GetItem(ID)
	if err != nil {
		return nil, exceptions.ErrInternalError
	}

	return event, nil
}

// Insert records an event in the database.
func (er *EventsRepository) Insert(events *domain.Events) error {
	return er.db.InsertItem(events)
}

package events

import (
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/core/domain/exceptions"
	"service-worker-sqs-dynamodb/dataproviders/awsdynamodb"
	"service-worker-sqs-dynamodb/dataproviders/mapper"
)

type IEventRepository interface {
	GetID(ID string) (*domain.Events, error)
	Insert(event *domain.Events) error
}

// EventRepository encapsulates all the data needed to the persistence in the event table.
type EventRepository struct {
	db *awsdynamodb.ClientDynamoDB
}

// NewEventRepository instance the connection to the dynamodb.
func NewEventRepository(db *awsdynamodb.ClientDynamoDB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

// GetID return the event by ID.
func (er *EventRepository) GetID(ID string) (*domain.Events, error) {
	event, err := er.db.GetItem(ID)
	if err != nil {
		return nil, exceptions.ErrInternalError
	}
	return mapper.ToDomain(event), nil
}

// Insert records an event in the database.
func (er *EventRepository) Insert(events *domain.Events) error {
	return er.db.InsertItem(mapper.ToEntity(events))
}

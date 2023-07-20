package mapper

import (
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/core/domain/entity"
)

func ToDomain(e *entity.Event) *domain.Events {
	return &domain.Events{
		ID:      e.ID,
		Message: e.Message,
		Date:    e.Date,
	}
}

func ToEntity(e *domain.Events) *entity.Event {
	return &entity.Event{
		ID:      e.ID,
		Message: e.Message,
		Date:    e.Date,
	}
}

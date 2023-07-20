package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/dataproviders/awssqs"
	"service-worker-sqs-dynamodb/dataproviders/consumer"
	repository "service-worker-sqs-dynamodb/dataproviders/repository/events"
)

// NewSQS define all usecases to instantiate SQS.
func NewSQS(logger *zap.SugaredLogger, config *Configuration, sessionaws *session.Session, repo repository.IEventsRepository) (domain.Source, error) {
	sqs, err := awssqs.NewSQSClient(sessionaws, config.SQSUrl, config.SQSMaxMessages, config.SQSVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("error awssqs.NewSQSClient: %w", err)
	}

	source, err := consumer.New(sqs, logger, config.SQSMaxMessages, repo)
	if err != nil {
		return nil, fmt.Errorf("error consumer.New: %w", err)
	}

	return source, nil
}

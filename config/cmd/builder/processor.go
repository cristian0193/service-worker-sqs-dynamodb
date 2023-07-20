package builder

import (
	"go.uber.org/zap"
	"service-worker-sqs-dynamo/core/domain"
	"service-worker-sqs-dynamo/dataproviders/processor"
)

// NewProcessor define all usecases to be instantiated Processor associated with the consumer.
func NewProcessor(logger *zap.SugaredLogger, source domain.Source) (*processor.Processor, error) {
	return processor.New(logger, source)
}

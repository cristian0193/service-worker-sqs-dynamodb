package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"service-worker-sqs-dynamodb/dataproviders/awsdynamodb"
)

// NewDynamodb define all usecases to instantiate DynamoDB.
func NewDynamodb(config *Configuration, sessionaws *session.Session) (*awsdynamodb.ClientDynamoDB, error) {
	db, err := awsdynamodb.NewDynamoDBClient(sessionaws, config.DynamoDBTable)
	if err != nil {
		return nil, fmt.Errorf("error awsdynamodb.NewDynamoDBClient: %w", err)
	}
	return db, nil
}

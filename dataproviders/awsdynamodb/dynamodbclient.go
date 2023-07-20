package awsdynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"service-worker-sqs-dynamodb/core/domain"
	"service-worker-sqs-dynamodb/dataproviders/utils"
)

// ClientDynamoDB represents DynamoDB client.
type ClientDynamoDB struct {
	api       *dynamodb.DynamoDB
	nameTable string
}

// NewDynamoDBClient instances of a Client to connect Dynamo with session as parameter.
func NewDynamoDBClient(sess *session.Session, nameTable string) (*ClientDynamoDB, error) {
	return &ClientDynamoDB{
		api:       dynamodb.New(sess),
		nameTable: nameTable,
	}, nil
}

// GetItem retrieves item from DynamoDB.
func (s *ClientDynamoDB) GetItem(id string) (*domain.Events, error) {
	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(s.nameTable),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	}
	res, err := s.api.GetItem(getItemInput)
	if err != nil {
		return nil, err
	}

	var events *domain.Events
	utils.Unmarshal(res.Item, &events)
	return events, nil
}

// InsertItem save item to DynamoDB.
func (s *ClientDynamoDB) InsertItem(event *domain.Events) error {
	item := map[string]*dynamodb.AttributeValue{
		"ID": {
			S: aws.String(event.ID),
		},
		"Message": {
			S: aws.String(event.Message),
		},
		"Date": {
			S: aws.String(event.Date),
		},
	}
	putItemInput := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(s.nameTable),
	}

	_, err := s.api.PutItem(putItemInput)
	if err != nil {
		return err
	}
	return nil
}

package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Unmarshal(item map[string]*dynamodb.AttributeValue, out interface{}) error {
	err := dynamodbattribute.UnmarshalMap(item, &out)
	if err != nil {
		return fmt.Errorf("error al convertir la interface: %v", err)
	}
	return nil
}

func Marshal(in interface{}) (map[string]*dynamodb.AttributeValue, error) {
	out, err := dynamodbattribute.MarshalMap(in)
	if err != nil {
		return nil, fmt.Errorf("error al convertir la interface: %v", err)
	}
	return out, nil
}

package utils

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func Unmarshal(item map[string]*dynamodb.AttributeValue, out interface{}) {
	err := dynamodbattribute.UnmarshalMap(item, &out)
	if err != nil {
		fmt.Println("Error al convertir el mapa a estructura:", err)
		return
	}
}

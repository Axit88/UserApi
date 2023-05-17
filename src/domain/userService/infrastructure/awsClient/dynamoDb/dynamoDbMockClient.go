package awsClient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbMockClient struct {
}

func (client DynamoDbMockClient) PushItemToDynamoDb(sess *session.Session, tableName string, id string, name string) (*dynamodb.PutItemOutput, error) {
	dummy := dynamodb.PutItemOutput{
		Attributes: map[string]*dynamodb.AttributeValue{
			"attribute1": {
				S: aws.String("Id"),
			},
			"attribute2": {
				N: aws.String("42"),
			},
		},
	}
	return &dummy, nil
}

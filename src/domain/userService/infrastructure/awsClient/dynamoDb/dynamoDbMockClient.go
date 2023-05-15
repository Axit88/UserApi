package awsClient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbMockClient struct {
}

func (client DynamoDbMockClient) PushItemToDynamoDb(sess *session.Session, tableName string, id string, name string) (*dynamodb.PutItemOutput, error) {
	return nil, nil
}

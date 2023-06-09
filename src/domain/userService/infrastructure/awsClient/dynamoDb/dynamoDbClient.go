package awsClient

import (
	"context"

	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamoDbImpl struct {
	logger *logger.LoggerImpl
}

func NewDynamoDbClient(l *logger.LoggerImpl) outgoing.DynamoDbClient {
	if constants.IsMock {
		return DynamoDbMockClient{}
	}

	return &DynamoDbImpl{
		logger: l,
	}
}

func (client DynamoDbImpl) PushItemToDynamoDb(sess *session.Session, tableName string, id string, name string) (*dynamodb.PutItemOutput, error) {
	input := dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"id":   {S: aws.String(id)},
			"name": {S: aws.String(name)},
		},
	}

	svc := dynamodb.New(sess)
	res, err := svc.PutItem(&input)
	if err != nil {
		client.logger.Errorf(context.Background(), "DynamoDb Client Failed", err)
		return nil, err
	}

	return res, nil
}

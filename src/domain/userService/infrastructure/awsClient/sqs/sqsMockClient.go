package awsClient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsMockClient struct {
}

func (client SqsMockClient) SendMessageToSqsQueue(sess *session.Session, queueUrl string, messageBody string) (*sqs.SendMessageOutput, error) {
	return nil, nil
}

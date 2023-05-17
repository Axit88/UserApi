package awsClient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsMockClient struct {
}

func (client SqsMockClient) SendMessageToSqsQueue(sess *session.Session, queueUrl string, messageBody string) (*sqs.SendMessageOutput, error) {
	res := sqs.SendMessageOutput{
		MessageId:        aws.String("message-id"),
		MD5OfMessageBody: aws.String("md5-hash"),
	}
	return &res, nil
}

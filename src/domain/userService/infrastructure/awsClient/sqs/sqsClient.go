package awsClient

import (
	"context"

	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type OutgoingSqs struct {
	logger *logger.LoggerImpl
}

func NewOutgoingSqsClient(l *logger.LoggerImpl) outgoing.SqsClient {
	return &OutgoingSqs{
		logger: l,
	}
}

func (client OutgoingSqs) SendMessageToSqsQueue(sess *session.Session, queueUrl string, messageBody string) (*sqs.SendMessageOutput, error) {

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(messageBody),
		QueueUrl:    aws.String(queueUrl),
	}

	svc := sqs.New(sess)
	result, err := svc.SendMessage(input)
	if err != nil {
		client.logger.Errorf(context.Background(), "Sqs Client Failed", err)
		return nil, err
	}

	return result, nil
}

package awsClient

import (
	"context"

	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type OutgoingKinesis struct {
	logger *logger.LoggerImpl
}

func NewOutgoingKinesisClient(l *logger.LoggerImpl) outgoing.KinesisClient {
	return &OutgoingKinesis{
		logger: l,
	}
}

func (client OutgoingKinesis) PushRecordToKinesis(sess *session.Session, kinesisStreamName string, data string, partitionKey string) (*kinesis.PutRecordOutput, error) {

	svc := kinesis.New(sess)
	res, err := svc.PutRecord(&kinesis.PutRecordInput{
		Data:         []byte(data),
		StreamName:   aws.String(kinesisStreamName),
		PartitionKey: aws.String(partitionKey),
	})

	if err != nil {
		client.logger.Errorf(context.Background(), "Kinesis Client Failed", err)
		return nil, err
	}

	return res, nil
}

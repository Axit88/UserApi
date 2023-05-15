package awsClient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisMockClient struct {
}

func (client KinesisMockClient) PushRecordToKinesis(sess *session.Session, kinesisStreamName string, data string, partitionKey string) (*kinesis.PutRecordOutput, error) {
	return nil, nil
}

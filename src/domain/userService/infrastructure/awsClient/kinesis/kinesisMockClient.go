package awsClient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

type KinesisMockClient struct {
}

func (client KinesisMockClient) PushRecordToKinesis(sess *session.Session, kinesisStreamName string, data string, partitionKey string) (*kinesis.PutRecordOutput, error) {
	res := kinesis.PutRecordOutput{
		EncryptionType: aws.String("NONE"),
		SequenceNumber: aws.String("1234567890"),
		ShardId:        aws.String("shard-001"),
	}
	return &res, nil
}

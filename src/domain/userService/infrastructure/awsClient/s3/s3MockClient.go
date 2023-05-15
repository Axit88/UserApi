package awsClient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3MockClient struct {
}

func (client S3MockClient) PutObjectInS3(sess *session.Session, bucketname string) (*s3.PutObjectOutput, error) {
	return nil, nil
}

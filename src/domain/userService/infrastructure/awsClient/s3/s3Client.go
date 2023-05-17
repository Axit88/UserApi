package awsClient

import (
	"bytes"
	"context"

	"github.com/Axit88/UserApi/src/constants"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Impl struct {
	logger *logger.LoggerImpl
}

func NewS3Client(l *logger.LoggerImpl) outgoing.S3Client {
	if constants.IsMock {
		return S3MockClient{}
	}

	return &S3Impl{
		logger: l,
	}
}

func (client S3Impl) PutObjectInS3(sess *session.Session, bucketname string) (*s3.PutObjectOutput, error) {

	svc := s3.New(sess)

	fileContent := []byte("This is a test file")
	res, err := svc.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(fileContent),
		Bucket: aws.String(bucketname),
		Key:    aws.String("test.txt"),
	})
	if err != nil {
		client.logger.Errorf(context.Background(), "S3 Client Failed", err)
		return nil, err
	}

	return res, nil
}

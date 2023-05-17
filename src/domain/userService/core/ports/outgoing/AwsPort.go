package outgoing

import (
	"database/sql"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type DynamoDbClient interface {
	PushItemToDynamoDb(sess *session.Session, tableName string, id string, name string) (*dynamodb.PutItemOutput, error)
}

type KinesisClient interface {
	PushRecordToKinesis(sess *session.Session, kinesisStreamName string, data string, partitionKey string) (*kinesis.PutRecordOutput, error)
}

type RdsClient interface {
	CreateDatabase(db *sql.DB, connection string, dbName string) error
	CreateTable(db *sql.DB, connection string, dbName string, tableName string) error
}

type S3Client interface {
	PutObjectInS3(sess *session.Session, bucketname string) (*s3.PutObjectOutput, error)
}

type RedisClient interface {
	RedisSetkey(key string, value string, expiryTime time.Duration) error
}

type SqsClient interface {
	SendMessageToSqsQueue(sess *session.Session, queueUrl string, messageBody string) (*sqs.SendMessageOutput, error)
}

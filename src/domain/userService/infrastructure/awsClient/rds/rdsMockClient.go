package awsClient

import (
	_ "github.com/go-sql-driver/mysql"
)

type RdsMockClient struct {
}

func (client RdsMockClient) CreateDatabase(dbName string) error {
	return nil
}

func (client RdsMockClient) CreateTable(dbName string, tableName string) error {
	return nil
}

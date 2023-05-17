package awsClient

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Axit88/UserApi/src/config"
	"github.com/Axit88/UserApi/src/constants"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
)

type RdsImpl struct {
	logger *logger.LoggerImpl
	db     *sql.DB
}

func NewRdsClient(l *logger.LoggerImpl) outgoing.RdsClient {
	if constants.IsMock {
		return RdsImpl{}
	}

	var cfn, _ = config.NewConfig()
	connection := fmt.Sprintf("%v:%v@tcp(%v)/%v", cfn.DbConfig.UserName, cfn.DbConfig.Password, cfn.DbConfig.Host, cfn.DbConfig.DatabaseName)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		l.Errorf(context.Background(), "db connection failure", err)
		log.Fatalf("db connection failure: %v", err)
	}

	return &RdsImpl{
		logger: l,
		db:     db,
	}
}

func (client RdsImpl) CreateDatabase(dbName string) error {

	_, err := client.db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		client.logger.Errorf(context.Background(), "Create Database Failed In RDS Client", err)
		return err
	}

	return nil
}

func (client RdsImpl) CreateTable(dbName string, tableName string) error {

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        age INT NOT NULL,
        PRIMARY KEY (id)
    )`, tableName)

	_, err := client.db.Exec(query)
	if err != nil {
		client.logger.Errorf(context.Background(), "Create Table Failed In RDS Client", err)
		return err
	}

	return nil
}

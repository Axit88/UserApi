package awsClient

import (
	"database/sql"
	"fmt"

	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/MindTickle/mt-go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
)

type OutgoingRds struct {
	logger *logger.LoggerImpl
}

func NewOutgoingRdsClient(l *logger.LoggerImpl) outgoing.RdsClient {
	return &OutgoingRds{
		logger: l,
	}
}

func (client OutgoingRds) CreateDatabase(db *sql.DB, connection string, dbName string) error {

	res, err := db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}

func (client OutgoingRds) CreateTable(db *sql.DB, connection string, dbName string, tableName string) error {

	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id INT NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
        age INT NOT NULL,
        PRIMARY KEY (id)
    )`, tableName)

	res, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}



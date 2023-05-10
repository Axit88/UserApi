package dbClient

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Axit88/UserApi/src/config"
	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
	"github.com/MindTickle/mt-go-logger/logger"
	_ "github.com/go-sql-driver/mysql"
)

type DbImpl struct {
	logger *logger.LoggerImpl
	db     *sql.DB
}

func NewDbClient(l *logger.LoggerImpl) (outgoing.DbPort, error) {
	if constants.IsMock {
		return DbMockClient{}, nil
	}

	var cfn, _ = config.NewConfig()
	connection := fmt.Sprintf("%v:%v@tcp(%v)/%v", cfn.DbConfig.UserName, cfn.DbConfig.Password, cfn.DbConfig.Host, cfn.DbConfig.DatabaseName)
	db, err := sql.Open("mysql", connection)
	if err != nil {
		l.Errorf(context.Background(), "db connection failure", err)
		log.Fatalf("db connection failure: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		l.Errorf(context.Background(), "db connection failure", err)
		log.Fatalf("db ping failure: %v", err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS USER (
			UserId VARCHAR(255) PRIMARY KEY,
			UserName VARCHAR(255)
		)
	`)
	if err != nil {
		l.Errorf(context.Background(), "db init failure", err)
		log.Fatalf("db init failure: %v", err)
	}

	return &DbImpl{db: db, logger: l}, nil
}

func (da DbImpl) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		da.logger.Errorf(context.Background(), "Db Close Failure", err)
		log.Fatalf("db close failure: %v", err)
	}
}

func (da DbImpl) Insert(input *model.User) error {
	var count int

	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", input.UserId).Scan(&count)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
		return err
	}

	if count > 0 {
		return fmt.Errorf("user %v already exist", input.UserId)
	}

	_, err = da.db.Exec("INSERT INTO USER (UserId, UserName) VALUES (?, ?)", input.UserId, input.UserName)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
	}

	return err
}

func (da DbImpl) Update(userId string, userName string) error {
	var count int

	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
		return err
	}

	if count == 0 {
		return fmt.Errorf("user %v not found", userId)
	}

	_, err = da.db.Exec("UPDATE USER SET UserName = ? WHERE UserId = ?", userName, userId)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
	}

	return err
}

func (da DbImpl) Select(userId string) (*model.User, error) {
	var count int

	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("user %v not found", userId)
	}

	queryResult, err := da.db.Query("SELECT * FROM USER WHERE UserId = ?", userId)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
		return nil, err
	}

	var id, name string
	for queryResult.Next() {
		err = queryResult.Scan(&id, &name)
		if err != nil {
			da.logger.Errorf(context.Background(), "Query Result Build Failed", err)
			return nil, err
		}
	}

	output := adapters.GetCreateUserRequest(id, name)
	return output, nil
}

func (da DbImpl) Delete(userId string) error {
	var count int

	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
		return err
	}

	if count == 0 {
		return fmt.Errorf("user %v not found", userId)
	}

	_, err = da.db.Exec("DELETE FROM USER WHERE UserId = ?", userId)
	if err != nil {
		da.logger.Errorf(context.Background(), "Query Execution Failed", err)
	}

	return err
}

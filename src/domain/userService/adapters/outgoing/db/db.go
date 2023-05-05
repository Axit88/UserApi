package db

import (
	"database/sql"
	"log"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"

	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) Insert(input *model.User) error {
	_, err := da.db.Exec("INSERT INTO USER (UserId, UserName) VALUES (?, ?)", input.UserId, input.UserName)
	if err != nil {
		return err
	}
	return nil
}

func (da Adapter) Update(userId string, userName string) error {
	_, err := da.db.Exec("UPDATE USER SET UserName = ? WHERE UserId = ?", userId, userName)
	return err
}

func (da Adapter) Select(userId string) (*model.User, error) {
	queryResult, err := da.db.Query("SELECT * FROM USER WHERE UserId = ?", userId)
	if err != nil {
		return nil, err
	}

	output := model.User{}
	for queryResult.Next() {
		err = queryResult.Scan(&output.UserId, &output.UserName)
		if err != nil {
			return nil, err
		}
	}
	return &output, nil
}

func (da Adapter) Delete(userId string) error {
	_, err := da.db.Exec("DELETE FROM USER WHERE UserId = ?", userId)
	return err
}

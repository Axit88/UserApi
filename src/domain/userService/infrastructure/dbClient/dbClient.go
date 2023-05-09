package dbClient

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Axit88/UserApi/src/config"
	"github.com/Axit88/UserApi/src/constants"
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	outgoing "github.com/Axit88/UserApi/src/domain/userService/core/ports/outgoing"
	"github.com/Axit88/UserApi/src/domain/userService/infrastructure/adapters"
	_ "github.com/go-sql-driver/mysql"
)

type DbImpl struct {
	db *sql.DB
}

func NewDbClient() (outgoing.DbPort, error) {
	if constants.IsMock {
		return DbMockClient{}, nil
	}

	var cfn, _ = config.NewConfig()
	connection := fmt.Sprintf("%v:%v@tcp(%v)/%v", cfn.DbConfig.UserName, cfn.DbConfig.Password, cfn.DbConfig.Host, cfn.DbConfig.DatabaseName)

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}

	// test db connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &DbImpl{db: db}, nil
}

func (da DbImpl) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da DbImpl) Insert(input *model.User) error {
	var count int
	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", input.UserId).Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return fmt.Errorf("User %v Already Exist", input.UserId)
	}
	_, err = da.db.Exec("INSERT INTO USER (UserId, UserName) VALUES (?, ?)", input.UserId, input.UserName)
	if err != nil {
		return err
	}
	return nil
}

func (da DbImpl) Update(userId string, userName string) error {
	var count int
	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("User %v Not Found", userId)
	}
	_, err = da.db.Exec("UPDATE USER SET UserName = ? WHERE UserId = ?", userName, userId)
	return err
}

func (da DbImpl) Select(userId string) (*model.User, error) {

	var count int
	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("User %v Not Found", userId)
	}

	queryResult, err := da.db.Query("SELECT * FROM USER WHERE UserId = ?", userId)
	if err != nil {
		return nil, err
	}

	var id,name string
	for queryResult.Next() {
		err = queryResult.Scan(&id, &name)
		if err != nil {
			return nil, err
		}
	}

	output := adapters.GetCreateUserRequest(id,name)
	return output, nil
}

func (da DbImpl) Delete(userId string) error {
	fmt.Println("no")
	var count int
	err := da.db.QueryRow("SELECT COUNT(*) FROM USER WHERE UserId=?", userId).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("User %v Not Found", userId)
	}
	_, err = da.db.Exec("DELETE FROM USER WHERE UserId = ?", userId)
	return err
}

package mtClient

import (
	"testing"

	"github.com/Axit88/UserApi/src/domain/userService/core/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateTable(t *testing.T) {

	var client = TickleDbStoreMockClient{}

	tableDetails := model.TickleDbEnvDetail{
		Env:       "test",
		TableName: "test",
		Namespace: "test",
	}

	res, err := client.CreateTable(tableDetails)
	
	assert.Nil(t, err)
	assert.Nil(t, res)
}

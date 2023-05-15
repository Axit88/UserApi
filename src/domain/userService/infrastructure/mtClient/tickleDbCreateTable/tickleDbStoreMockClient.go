package main

import (
	"github.com/Axit88/UserApi/src/domain/userService/core/model"
)

type TickleDbStoreMockClient struct {
}

func (client TickleDbStoreMockClient) CreateTable(dbDetail model.TickleDbEnvDetail) error {
	return nil
}

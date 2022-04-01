package database

import (
	"testing"
)

func TestConnectMySQL(t *testing.T) {

	ConnectDB("root:root@tcp(127.0.0.1:3306)/test-wallet-api?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := DB.DB()
	if err != nil {
		t.Error(err)
	}
	err = db.Ping()
	if err != nil {
		t.Error(err)
	}
}

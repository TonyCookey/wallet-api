package database

import (
	"os"
	"testing"
)

func TestConnectMySQL(t *testing.T) {
	ConnectDB(os.Getenv("DATABASE_URL"))

	db, err := DB.DB()
	if err != nil {
		t.Error(err)
	}
	err = db.Ping()
	if err != nil {
		t.Error(err)
	}
}

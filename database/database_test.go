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

//func TestAutoMigration(t *testing.T) {
//	ConnectDB(os.Getenv("DATABASE_URL"))
//
//
//	type test struct {
//		Name string
//	}
//	AutoMigration(test{})
//	_, tableCheck := DB.Comm
//	if tableCheck == nil {
//		t.Error("Table is not created")
//	}
//}

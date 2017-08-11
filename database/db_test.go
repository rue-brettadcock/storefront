package database

import (
	"testing"
)

func TestOpenDatabaseConnection(t *testing.T) {
	mydb := MyDb{db: nil}

	mydb.openDatabaseConnection()

	if mydb.db.Ping() != nil {
		t.Error("Could not connect to the database")
	}
}

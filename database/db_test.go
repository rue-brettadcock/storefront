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

func TestPrint(t *testing.T) {
	mydb := New()
	expected := ""
	mydb.db.Exec(
		"INSERT INTO products(id, name, vendor, quantity) VALUES(?,?,?,?)",
		222, "n", "v", 20)
	if mydb.Print() != expected {
		t.Error("Not printing correctly")
	}
}

package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//MyDb is a struct to restrict access to the db
type MyDb struct {
	db *sql.DB
}

//New initializes a pointer to a sql database
func New() *MyDb {
	m := MyDb{db: nil}
	m.openDatabaseConnection()

	return &m
}

func (s *MyDb) openDatabaseConnection() {
	var err error
	s.db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/productInfo")
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

//Delete removes an entry based on id from the products table in productInfo db
func (s *MyDb) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id=?", id)
	return err
}

//Insert puts given product information into the products table in the db
func (s *MyDb) Insert(id int, name string, vendor string, quantity int) error {
	_, err := s.db.Exec("INSERT INTO products(id, name, vendor, quantity) VALUES(?, ?, ?, ?)",
		id, name, vendor, quantity)
	return err
}

//Update changes the products quantity
func (s *MyDb) Update(id, quantity int) error {
	_, err := s.db.Exec("UPDATE products SET quantity=? WHERE id=?", quantity, id)
	return err
}

//Get returns the product info for a given id
func (s *MyDb) Get(id int) string {
	res, err := s.db.Query("SELECT * FROM products WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
	return printRows(res)
}

//Print prints stuff from database
func (s *MyDb) Print() string {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}

	result := printRows(rows)
	return result
}

func printRows(rows *sql.Rows) string {
	result := ""

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			result += columns[i] + ": " + value + "\n"
			//fmt.Println(columns[i], ": ", value)
		}
		result += "-----------------------------------\n"
		//fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

//MyDb is a struct to restrict access to the db
type MyDb struct {
	db *sql.DB
	mu sync.Mutex
}

//New initializes a pointer to a sql database
func New() *MyDb {
	m := MyDb{db: nil}
	m.openDatabaseConnection()

	return &m
}

//Close ends the connection to the database
func (s *MyDb) Close() {
	s.db.Close()
}

func (s *MyDb) openDatabaseConnection() {
	var err error
	s.mu.Lock()
	s.db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/productInfo")
	if err != nil {
		log.Fatal(err)
	}
	err = s.db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	s.mu.Unlock()
}

//Delete removes an entry based on id from the products table in productInfo db
func (s *MyDb) Delete(id int) error {
	s.mu.Lock()
	_, err := s.db.Exec("DELETE FROM products WHERE id=?", id)
	s.mu.Unlock()
	return err
}

//Insert puts given product information into the products table in the db
func (s *MyDb) Insert(id int, name string, vendor string, quantity int) error {
	s.mu.Lock()
	_, err := s.db.Exec("INSERT INTO products(id, name, vendor, quantity) VALUES(?, ?, ?, ?)",
		id, name, vendor, quantity)
	s.mu.Unlock()
	return err
}

//Update changes the products quantity
func (s *MyDb) Update(id, quantity int) error {
	s.mu.Lock()
	_, err := s.db.Exec("UPDATE products SET quantity=? WHERE id=?", quantity, id)
	s.mu.Unlock()
	return err
}

//Get returns the product info for a given id
func (s *MyDb) Get(id int) string {
	s.mu.Lock()
	res, err := s.db.Query("SELECT * FROM products WHERE id=?", id)
	s.mu.Unlock()
	if err != nil {
		return ""
	}
	return printRows(res)
}

//Print prints product information from database
func (s *MyDb) Print() string {
	s.mu.Lock()
	rows, err := s.db.Query("SELECT * FROM products")
	s.mu.Unlock()
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

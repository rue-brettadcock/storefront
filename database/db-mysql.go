package database

import (
	"database/sql"
	"encoding/json"
	"log"

	//drivers for mysql connection
	_ "github.com/go-sql-driver/mysql"
)

//SQLdb is a struct to restrict access to the db
type SQLdb struct {
	db *sql.DB
}

func openDatabaseConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/productInfo")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//Delete removes an entry based on id from the products table in productInfo db
func (s *SQLdb) Delete(id string) error {
	_, err := s.db.Exec("DELETE FROM products WHERE id=?", id)
	return err
}

//Insert puts given product information into the products table in the db
func (s *SQLdb) Insert(id string, name string, vendor string, quantity int) error {
	_, err := s.db.Exec("INSERT INTO products(id, name, vendor, quantity) VALUES(?, ?, ?, ?)",
		id, name, vendor, quantity)
	return err
}

//Update changes the products quantity
func (s *SQLdb) Update(id string, quantity int) error {
	_, err := s.db.Exec("UPDATE products SET quantity=? WHERE id=?", quantity, id)
	return err
}

//Get returns the product info for a given id
func (s *SQLdb) Get(id string) string {
	res, err := s.buildJSON("SELECT * FROM products WHERE id=" + id)
	if err != nil {
		return ""
	}
	return res
}

//Print prints product information from database
func (s *SQLdb) Print() string {
	res, err := s.buildJSON("SELECT * FROM products")
	if err != nil {
		return ""
	}
	return res
}

func (s *SQLdb) buildJSON(queryStr string) (string, error) {
	rows, err := s.db.Query(queryStr)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}

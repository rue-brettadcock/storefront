package database

import "fmt"

type sku struct {
	id     int
	name   string
	vendor string
	amt    int
}

//MemDb is a struct to restrict access to the db
type MemDb struct {
	db []sku
}

var m MemDb

func newConnection() []sku {
	return m.db
}

//Insert puts given product information into a sku in the slice
func (m *MemDb) Insert(id int, name string, vendor string, amt int) error {
	m.db = append(m.db, sku{id, name, vendor, amt})
	return nil
}

//Get returns the product info for a given id
func (m *MemDb) Get(id int) string {
	for _, s := range m.db {
		if s.id == id {
			var toPrint []sku
			toPrint = append(toPrint, s)
			return buildJSON(toPrint)
		}
	}
	return "[]"
}

//Print prints product information from database
func (m *MemDb) Print() string {
	return buildJSON(m.db)
}

//Update changes the products quantity
func (m *MemDb) Update(id, amt int) error {
	for _, s := range m.db {
		if s.id == id {
			s.amt = amt
			break
		}
	}
	return nil
}

//Delete removes the sku with the matching id
func (m *MemDb) Delete(id int) error {
	for i, s := range m.db {
		if s.id == id {
			m.db = append(m.db[:i], m.db[i+1:]...)
			break
		}
	}
	return nil
}

func buildJSON(list []sku) string {
	if len(list) == 0 {
		return "[]"
	}

	res := ""
	for _, s := range list {
		res += fmt.Sprintf("{\"id\":\"%v\",\"name\":\"%v\",\"vendor\":\"%v\",\"quantity\":\"%v\"},", s.id, s.name, s.vendor, s.amt)
	}
	res = "[" + res[:len(res)-1] + "]"

	return res
}

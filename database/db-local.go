package database

import (
	"encoding/json"
)

type sku struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
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
func (m *MemDb) Insert(id string, name string, vendor string, amt int) error {
	m.db = append(m.db, sku{id, name, vendor, amt})
	return nil
}

//Get returns the product info for a given id
func (m *MemDb) Get(id string) string {
	for _, s := range m.db {
		if s.ID == id {
			res, _ := json.Marshal(&s)
			return string(res)
		}
	}
	return "[]"
}

//Print prints product information from database
func (m *MemDb) Print() string {
	res, _ := json.Marshal(m.db)
	if string(res) == "null" {
		return "[]"
	}
	return string(res)
}

//Update changes the products quantity
func (m *MemDb) Update(id string, amt int) error {
	for _, s := range m.db {
		if s.ID == id {
			s.Quantity = amt
			break
		}
	}
	return nil
}

//Delete removes the sku with the matching id
func (m *MemDb) Delete(id string) error {
	for i, s := range m.db {
		if s.ID == id {
			m.db = append(m.db[:i], m.db[i+1:]...)
			break
		}
	}
	return nil
}

package database

import (
	"encoding/json"
	"sync"

	errors "github.com/pkg/errors"
)

//MemDb is a struct to restrict access to the db
type MemDb struct {
	db []sku
	mu sync.Mutex
}

var m MemDb

func newConnection() []sku {
	return m.db

}

//Insert puts given product information into a sku in the slice
func (m *MemDb) Insert(id string, name string, vendor string, amt int) error {
	m.mu.Lock()
	m.db = append(m.db, sku{id, name, vendor, amt})
	m.mu.Unlock()
	return nil
}

//Get returns the product info for a given id
func (m *MemDb) Get(id string) string {
	result := "[]"
	m.mu.Lock()
	for _, s := range m.db {
		if s.ID == id {
			json, _ := json.Marshal(&s)
			result = string(json)
			break
		}
	}
	m.mu.Unlock()
	return result
}

//Print prints product information from database
func (m *MemDb) Print() string {
	m.mu.Lock()
	res, _ := json.Marshal(m.db)
	m.mu.Unlock()
	if string(res) == "null" {
		return "[]"
	}
	return string(res)
}

//Update changes the products quantity
func (m *MemDb) Update(id string, amt int) error {
	position := -1
	m.mu.Lock()
	for i, s := range m.db {
		if s.ID == id {
			position = i
			break
		}
	}
	if position == -1 {
		return errors.New("Failed to update sku")
	}
	m.db[position].Quantity = amt
	m.mu.Unlock()
	return nil
}

//Delete removes the sku with the matching id
func (m *MemDb) Delete(id string) error {
	m.mu.Lock()
	for i, s := range m.db {
		if s.ID == id {
			m.db = append(m.db[:i], m.db[i+1:]...)
			break
		}
	}
	m.mu.Unlock()
	return nil
}

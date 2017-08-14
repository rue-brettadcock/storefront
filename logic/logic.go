package logic

import (
	"errors"
	"log"

	"github.com/rue-brettadcock/storefront/database"
)

//Logic contains a pointer to a database instance
type Logic struct {
	mydb *database.MyDb
}

//New creates a new logic pointer to the database layer
func New() *Logic {
	l := Logic{mydb: database.New()}
	return &l
}

//AddProductSKU validates product info and Inserts into the db
func (l *Logic) AddProductSKU(id int, name string, vendor string, quantity int) error {
	if l.mydb.Get(id) != "" {
		return errors.New("Product id already exists")
	}
	if quantity < 1 {
		return errors.New("Quantity must be at least 1")
	}
	if id < 0 {
		return errors.New("ID must be positive")
	}

	err := l.mydb.Insert(id, name, vendor, quantity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//UpdateProductQuantity updates quantity for a given id
func (l *Logic) UpdateProductQuantity(id, quantity int) error {
	if l.mydb.Get(id) == "" {
		return errors.New("Product id doesn't exist")
	}

	err := l.mydb.Update(id, quantity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//DeleteID removes all product information for a given id
func (l *Logic) DeleteID(id int) error {
	err := l.mydb.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//PrintAllProductInfo returns all product SKUs
func (l *Logic) PrintAllProductInfo(id int) string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *Logic) GetProductInfo(id int) string {
	return l.mydb.Get(id)
}

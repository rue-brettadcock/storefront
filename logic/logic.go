package logic

import (
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
func (l *Logic) AddProductSKU(id int, name string, vendor string, quantity int) string {
	if l.mydb.Get(id) != "" {
		return "Product id already exists"
	}
	if quantity < 1 {
		return "Quantity must be at least 1"
	}
	if id < 0 {
		return "ID must be positive"
	}

	err := l.mydb.Insert(id, name, vendor, quantity)
	if err != nil {
		log.Fatal(err)
	}
	return "Product successfully added to database"
}

//UpdateProductQuantity updates quantity for a given id
func (l *Logic) UpdateProductQuantity(id, quantity int) string {
	if l.mydb.Get(id) == "" {
		return "Product id doesn't exist"
	}

	err := l.mydb.Update(id, quantity)
	if err != nil {
		log.Fatal(err)
	}
	return "SKU successfully updated"
}

//DeleteID removes all product information for a given id
func (l *Logic) DeleteID(id int) string {
	if l.mydb.Get(id) == "" {
		return "Product id doesn't exist"
	}
	err := l.mydb.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	return "Product successfully deleted"
}

//PrintAllProductInfo returns all product SKUs
func (l *Logic) PrintAllProductInfo(id int) string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *Logic) GetProductInfo(id int) string {
	return l.mydb.Get(id)
}

package logic

import (
	"log"

	"github.com/rue-brettadcock/storefront/database"
)

//Logic contains a pointer to a database instance
type logic struct {
	mydb database.SKUDataAccess
}

//Logic explicit interface for ioc
type Logic interface {
	AddProductSKU(SKU) string
	UpdateProductQuantity(SKU) string
	DeleteID(SKU) string
	PrintAllProductInfo() string
	GetProductInfo(SKU) string
}

//SKU for holding product information
type SKU struct {
	ID       int
	Name     string
	Vendor   string
	Quantity int
}

//New creates a new logic pointer to the database layer
func New() Logic {
	l := logic{mydb: database.New()}
	return &l
}

//AddProductSKU validates product info and Inserts into the db
func (l *logic) AddProductSKU(sku SKU) string {
	if l.mydb.Get(sku.ID) != "" {
		return "Product id already exists"
	}
	if sku.Quantity < 1 {
		return "Quantity must be at least 1"
	}
	if sku.ID < 0 {
		return "ID must be positive"
	}

	err := l.mydb.Insert(sku.ID, sku.Name, sku.Vendor, sku.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return "Product successfully added to database"
}

//UpdateProductQuantity updates quantity for a given id
func (l *logic) UpdateProductQuantity(sku SKU) string {
	if l.mydb.Get(sku.ID) == "" {
		return "Product id doesn't exist"
	}

	err := l.mydb.Update(sku.ID, sku.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return "SKU successfully updated"
}

//DeleteID removes all product information for a given id
func (l *logic) DeleteID(sku SKU) string {
	if l.mydb.Get(sku.ID) == "" {
		return "Product id doesn't exist"
	}
	err := l.mydb.Delete(sku.ID)
	if err != nil {
		log.Fatal(err)
	}
	return "Product successfully deleted"
}

//PrintAllProductInfo returns all product SKUs
func (l *logic) PrintAllProductInfo() string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *logic) GetProductInfo(sku SKU) string {
	if l.mydb.Get(sku.ID) == "" {
		return "Product id doesn't exist"
	}
	return l.mydb.Get(sku.ID)
}

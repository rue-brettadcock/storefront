package logic

import (
	"errors"
	"log"

	"github.com/rue-brettadcock/storefront/database"
)

//Logic contains a pointer to a database instance
type logic struct {
	mydb database.SKUDataAccess
}

//Logic explicit interface for ioc
type Logic interface {
	AddProductSKU(SKU) error
	UpdateProductQuantity(SKU) error
	DeleteID(SKU) error
	PrintAllProductInfo() string
	GetProductInfo(SKU) (string, error)
}

//SKU for holding product information
type SKU struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

//New creates a new logic pointer to the database layer
func New() Logic {
	l := logic{mydb: database.NewInMemoryDB()}
	return &l
}

//AddProductSKU validates product info and Inserts into the db
func (l *logic) AddProductSKU(sku SKU) error {
	if l.mydb.Get(sku.ID) != "[]" {
		return errors.New("Product id already exists")
	}
	if sku.Quantity < 1 {
		return errors.New("Quantity must be at least 1")
	}
	if sku.ID < 0 {
		return errors.New("ID must be positive")
	}

	err := l.mydb.Insert(sku.ID, sku.Name, sku.Vendor, sku.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//UpdateProductQuantity updates quantity for a given id
func (l *logic) UpdateProductQuantity(sku SKU) error {
	if l.mydb.Get(sku.ID) == "[]" {
		return errors.New("Product id doesn't exist")
	}

	err := l.mydb.Update(sku.ID, sku.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//DeleteID removes all product information for a given id
func (l *logic) DeleteID(sku SKU) error {
	if l.mydb.Get(sku.ID) == "[]" {
		return errors.New("Product id doesn't exist")
	}
	err := l.mydb.Delete(sku.ID)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

//PrintAllProductInfo returns all product SKUs
func (l *logic) PrintAllProductInfo() string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *logic) GetProductInfo(sku SKU) (string, error) {
	info := l.mydb.Get(sku.ID)
	if info == "[]" {
		return info, errors.New("Product id doesn't exist")
	}
	return info, nil
}

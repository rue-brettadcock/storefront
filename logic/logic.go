package logic

import (
	"log"
	"strconv"

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
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

//New creates a new logic pointer to the database layer
func New() Logic {
	l := logic{mydb: database.New()}
	return &l
}

//AddProductSKU validates product info and Inserts into the db
func (l *logic) AddProductSKU(sku SKU) string {
	id, _ := strconv.Atoi(sku.ID)
	quant, _ := strconv.Atoi(sku.Quantity)
	if l.mydb.Get(id) != "" {
		return "Product id already exists"
	}
	if quant < 1 {
		return "Quantity must be at least 1"
	}
	if id < 0 {
		return "ID must be positive"
	}

	err := l.mydb.Insert(id, sku.Name, sku.Vendor, quant)
	if err != nil {
		log.Fatal(err)
	}
	return "Product successfully added to database"
}

//UpdateProductQuantity updates quantity for a given id
func (l *logic) UpdateProductQuantity(sku SKU) string {
	id, _ := strconv.Atoi(sku.ID)
	quant, _ := strconv.Atoi(sku.Quantity)
	if l.mydb.Get(id) == "" {
		return "Product id doesn't exist"
	}

	err := l.mydb.Update(id, quant)
	if err != nil {
		log.Fatal(err)
	}
	return "SKU successfully updated"
}

//DeleteID removes all product information for a given id
func (l *logic) DeleteID(sku SKU) string {
	id, _ := strconv.Atoi(sku.ID)
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
func (l *logic) PrintAllProductInfo() string {
	return l.mydb.Print()
}

//GetProductInfo returns product details for given id
func (l *logic) GetProductInfo(sku SKU) string {
	id, _ := strconv.Atoi(sku.ID)
	if l.mydb.Get(id) == "" {
		return "Product id doesn't exist"
	}
	return l.mydb.Get(id)
}

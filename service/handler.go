package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rue-brettadcock/storefront/logic"
)

//Presentation for isolating access to logic layer
type Presentation struct {
	logic logic.Logic
}

func (p *Presentation) addSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	json.NewDecoder(req.Body).Decode(&sku)

	err := p.logic.AddProductSKU(sku)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusCreated)
}

func (p *Presentation) printSKUs(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	msg := p.logic.PrintAllProductInfo()
	fmt.Fprintf(res, "%v\n", msg)
	if msg == "[]" {
		res.WriteHeader(http.StatusNoContent)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (p *Presentation) getSKU(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var sku logic.SKU
	params := mux.Vars(req)
	sku.ID = params["id"]
	msg, err := p.logic.GetProductInfo(sku)
	fmt.Fprintf(res, "%v\n", msg)
	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (p *Presentation) updateSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	json.NewDecoder(req.Body).Decode(&sku)

	err := p.logic.UpdateProductQuantity(sku)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (p *Presentation) deleteSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	params := mux.Vars(req)
	sku.ID = params["id"]

	err := p.logic.DeleteID(sku)
	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		return
	}
	res.WriteHeader(http.StatusOK)
}

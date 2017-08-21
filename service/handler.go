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
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.AddProductSKU(sku)
	fmt.Fprintf(res, "%v\n", msg)
}

func (p *Presentation) printSKUs(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.PrintAllProductInfo()
	fmt.Fprintf(res, "%v\n", msg)
}

func (p *Presentation) getSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	params := mux.Vars(req)
	sku.ID = params["id"]
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.GetProductInfo(sku)
	fmt.Fprintf(res, "%v\n", msg)
}

func (p *Presentation) updateSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	json.NewDecoder(req.Body).Decode(&sku)
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.UpdateProductQuantity(sku)
	fmt.Fprintf(res, "%v\n", msg)
}

func (p *Presentation) deleteSKU(res http.ResponseWriter, req *http.Request) {
	var sku logic.SKU
	json.NewDecoder(req.Body).Decode(&sku)
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.DeleteID(sku)
	fmt.Fprintf(res, "%v\n", msg)
}

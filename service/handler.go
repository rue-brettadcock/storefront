package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/rue-brettadcock/storefront/logic"

	"goji.io/pat"
)

//Presentation for isolating access to logic layer
type Presentation struct {
	logic logic.Logic
}

func (p *Presentation) addSKU(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(pat.Param(req, "id"))
	name := pat.Param(req, "name")
	vendor := pat.Param(req, "vendor")
	quantity, _ := strconv.Atoi(pat.Param(req, "quantity"))

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.AddProductSKU(id, name, vendor, quantity)
	fmt.Fprintf(res, "%s\n", msg)
}

func (p *Presentation) updateSKU(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(pat.Param(req, "id"))
	quantity, _ := strconv.Atoi(pat.Param(req, "quantity"))

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	msg := p.logic.UpdateProductQuantity(id, quantity)
	fmt.Fprintf(res, "%s\n", msg)
}

func (p *Presentation) deleteSKU(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(pat.Param(req, "id"))

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	msg := p.logic.DeleteID(id)
	fmt.Fprintf(res, "%s/n", msg)
}

func (p *Presentation) printSKUs(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	output := p.logic.PrintAllProductInfo()
	fmt.Fprintf(res, "%s", output)
}

func (p *Presentation) getSKU(res http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(pat.Param(req, "id"))

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	output := p.logic.GetProductInfo(id)
	fmt.Fprintf(res, "%s", output)
}

package service

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rue-brettadcock/storefront/logic"
)

//Presentation for isolating access to logic layer
type Presentation struct {
	logic logic.Logic
}

func (p *Presentation) addSKU(res http.ResponseWriter, req *http.Request) {
	u := formatPath(req.RequestURI)
	values, _ := url.ParseQuery(u)
	sku := buildSKU(values)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)

	msg := p.logic.AddProductSKU(sku)
	fmt.Fprintf(res, "%s\n", msg)
}

func (p *Presentation) updateSKU(res http.ResponseWriter, req *http.Request) {
	u := formatPath(req.RequestURI)
	values, _ := url.ParseQuery(u)
	sku := buildSKU(values)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	msg := p.logic.UpdateProductQuantity(sku)
	fmt.Fprintf(res, "%s\n", msg)
}

func (p *Presentation) deleteSKU(res http.ResponseWriter, req *http.Request) {
	u := formatPath(req.RequestURI)
	values, _ := url.ParseQuery(u)
	sku := buildSKU(values)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	msg := p.logic.DeleteID(sku)
	fmt.Fprintf(res, "%s/n", msg)
}

func (p *Presentation) printSKUs(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	output := p.logic.PrintAllProductInfo()
	fmt.Fprintf(res, "%s", output)
}

func (p *Presentation) getSKU(res http.ResponseWriter, req *http.Request) {
	u := formatPath(req.RequestURI)
	values, _ := url.ParseQuery(u)
	sku := buildSKU(values)

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
	output := p.logic.GetProductInfo(sku)
	fmt.Fprintf(res, "%s", output)
}

func buildSKU(vals url.Values) logic.SKU {
	var sku logic.SKU
	sku.ID, _ = strconv.Atoi(vals.Get("id"))
	sku.Name = vals.Get("name")
	sku.Vendor = vals.Get("vend")
	sku.Quantity, _ = strconv.Atoi(vals.Get("amt"))

	return sku
}

func formatPath(uri string) string {
	rmPrefix, result := "", ""
	numSlash := 0

	for i, r := range uri {
		c := string(r)
		if c == "/" {
			numSlash++
		}
		if numSlash == 2 {
			rmPrefix = uri[i+1:]
			break
		}
	}
	for _, r := range rmPrefix {
		c := string(r)
		if c == "?" {
			result += ";"
		} else {
			result += c
		}
	}
	return result
}

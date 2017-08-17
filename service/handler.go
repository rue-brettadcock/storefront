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

func (p *Presentation) handleHTTP(res http.ResponseWriter, req *http.Request) {
	u, prefix := formatPath(req.RequestURI)
	fmt.Printf("prefix: " + prefix + "\nURL: " + u)
	values, _ := url.ParseQuery(u)
	sku := buildSKU(values)

	res.Header().Set("Content-Type", "text/plain")

	var msg string

	switch prefix {
	case "/addSKU":
		msg = p.logic.AddProductSKU(sku)
	case "/updateSKU":
		msg = p.logic.UpdateProductQuantity(sku)
	case "/deleteSKU":
		msg = p.logic.DeleteID(sku)
	case "/printSKUs":
		msg = p.logic.PrintAllProductInfo()
	case "/getSKU":
		msg = p.logic.GetProductInfo(sku)
	default:
		res.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprintf(res, "%s\n", msg)
}

func buildSKU(vals url.Values) logic.SKU {
	var sku logic.SKU
	sku.ID, _ = strconv.Atoi(vals.Get("id"))
	sku.Name = vals.Get("name")
	sku.Vendor = vals.Get("vend")
	sku.Quantity, _ = strconv.Atoi(vals.Get("amt"))

	return sku
}

func formatPath(uri string) (string, string) {
	prefix, suffix, result := "", "", ""
	numSlash := 0

	for i, r := range uri {
		c := string(r)
		if c == "/" {
			numSlash++
		}
		if numSlash == 2 {
			prefix = uri[:i]
			suffix = uri[i+1:]
			break
		}
	}
	for _, r := range suffix {
		c := string(r)
		if c == "?" {
			result += ";"
		} else {
			result += c
		}
	}
	return result, prefix
}

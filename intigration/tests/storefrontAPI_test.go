package tests

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"

	"github.com/emicklei/forest"
	uuid "github.com/satori/go.uuid"
)

//local struct modeled after logic.SKU
//used to ease encoding/decoding of req.Body json values
type sku struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

var sf = forest.NewClient("http://localhost:8080", new(http.Client))

// func TestPrintSKUs_NoExistingSKUs_NoContent(t *testing.T) {
// 	printSKU := sf.GET(t, forest.Path("/products"))
// 	if forest.ExpectStatus(t, printSKU, http.StatusNoContent) != true {
// 		t.Errorf("Actual: %v\nExpected: %v", printSKU.StatusCode, http.StatusNoContent)
// 	}
// }

func TestAddSKU_HappyPath_Success(t *testing.T) {
	sample := sku{ID: 123, Name: "polo", Vendor: "RL", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	addSKU := sf.POST(t, forest.Path("/products").Body(string(productInfo)))
	if forest.ExpectStatus(t, addSKU, http.StatusCreated) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusCreated)
	}

	getSKU := sf.GET(t, forest.Path("/products/123"))
	if forest.ExpectStatus(t, getSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusOK)
	}

	body, _ := ioutil.ReadAll(getSKU.Body)
	var actual sku
	json.Unmarshal(body, &actual)
	if actual != sample {
		t.Errorf("Actual: %v\nExpected: %v", actual, sample)
	}

}

func TestAddSKU_addSameItemTwice_BadRequest(t *testing.T) {
	sample := sku{ID: 345, Name: "pen", Vendor: "bic", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	addSKU := sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	if forest.ExpectStatus(t, addSKU, http.StatusCreated) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusCreated)
	}
	addSKU = sf.POST(t, forest.Path("/products").Body(string(productInfo)))
	if forest.ExpectStatus(t, addSKU, http.StatusBadRequest) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusBadRequest)
	}
}

func TestDeleteID_idDoesntExist_NoContent(t *testing.T) {
	delSKU := sf.DELETE(t, forest.Path("/products/1"))
	if forest.ExpectStatus(t, delSKU, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU.StatusCode, http.StatusNoContent)
	}
}

func TestDeleteID_idExists_Success(t *testing.T) {
	sample := sku{ID: 678, Name: "watch", Vendor: "breitling", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	delSKU := sf.DELETE(t, forest.Path("/products/678"))
	if forest.ExpectStatus(t, delSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU.StatusCode, http.StatusOK)
	}
}

func TestUpdateSKU_productExists_Success(t *testing.T) {
	sample := sku{ID: 678, Name: "watch", Vendor: "breitling", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	b := `{"id":678,"quantity":1000}`
	updateSKU := sf.PUT(t, forest.Path("/products").Body(b))
	if forest.ExpectStatus(t, updateSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", updateSKU.StatusCode, http.StatusOK)
	}

	getSKU := sf.GET(t, forest.Path("/products/678"))
	body, _ := ioutil.ReadAll(getSKU.Body)
	var actual sku
	json.Unmarshal(body, &actual)
	if actual != sample {
		t.Errorf("Actual: %v\nExpected: %v", actual, sample)
	}
}

func TestUpdateSKU_productDoesntExists_BadReqest(t *testing.T) {
	b := `{"id":890,"quantity":1000}`
	updateSKU := sf.PUT(t, forest.Path("/products").Body(b))
	if forest.ExpectStatus(t, updateSKU, http.StatusBadRequest) != true {
		t.Errorf("Actual: %v\nExpected: %v", updateSKU.StatusCode, http.StatusBadRequest)
	}
}

func TestGetSKU_idExists_Success(t *testing.T) {
	sample := sku{ID: 456, Name: "clock", Vendor: "atomic", Quantity: 120}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	getSKU := sf.GET(t, forest.Path("/products/456"))
	if forest.ExpectStatus(t, getSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusOK)
	}

	body, _ := ioutil.ReadAll(getSKU.Body)
	var actual sku
	json.Unmarshal(body, &actual)
	if actual != sample {
		t.Errorf("Actual: %v\nExpected: %v", actual, sample)
	}
}

func TestGetSKU_idDoesntExists_BadRequest(t *testing.T) {
	getSKU := sf.GET(t, forest.Path("/products/404"))
	if forest.ExpectStatus(t, getSKU, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusNoContent)
	}
}

func newIntUUID() *big.Int {
	u := uuid.NewV1()
	var i big.Int
	res, _ := i.SetString(strings.Replace(u.String(), "-", "", 4), 16)
	return res
}

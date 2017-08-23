package tests

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/emicklei/forest"
	"github.com/rue-brettadcock/storefront/logic"
)

var sf = forest.NewClient("http://localhost:8080", new(http.Client))

func TestAddSKU_HappyPath_Success(t *testing.T) {
	sample := logic.SKU{ID: 123, Name: "polo", Vendor: "RL", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	addSKU := sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	if forest.ExpectStatus(t, addSKU, http.StatusCreated) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusCreated)
	}

	getSKU := sf.GET(t, forest.Path("/products/123"))

	if forest.ExpectStatus(t, getSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusOK)
	}
	if forest.ExpectJSONDocument(t, getSKU, &productInfo) != true {
		t.Error("i broke it")
	}

}

// func TestAddSKU_HappyPath_Success_1(t *testing.T) {
// 	product := logic.SKU{ID: 21, Name: "polo", Vendor: "RL", Quantity: 10}
// 	productJSON, _ := json.Marshal(&product)
// 	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 	rr := httptest.NewRecorder()
// 	handler := http.Handler()

// 	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 	// directly and pass in our Request and ResponseRecorder.
// 	handler.ServeHTTP(rr, req)

// 	// Check the status code is what we expect.
// 	if status := rr.Code; status != http.StatusCreated {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusCreated)
// 	}

// 	// Check the response body is what we expect.
// 	expected := ``
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}

// }

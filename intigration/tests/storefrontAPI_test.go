package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"

	"github.com/emicklei/forest"
	"github.com/rue-brettadcock/storefront/intigration/client"
	uuid "github.com/satori/go.uuid"
)

//local struct modeled after logic.SKU
//used to ease encoding/decoding of req.Body json values
type sku struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

var sf = forest.NewClient("http://localhost:8080", new(http.Client))

func initServ() {

	t := &testing.T{}
	server := forest.NewClient("http://localhost:8080", new(http.Client))
	//uid := newUUID()
	sku1 := sku{ID: "12", Name: "polo", Vendor: "RL", Quantity: 25}
	productInfo, _ := json.Marshal(&sku1)
	server.POST(t, forest.Path("/products").Body(string(productInfo)))

}

func closeServ() {
	t := &testing.T{}
	server := forest.NewClient("http://localhost:8080", new(http.Client))
	server.DELETE(t, forest.Path("/products/12"))

}

func add1(num int) int {
	return num + 1
}

func BenchmarkRaceTest_updateSKU(b *testing.B) {
	initServ()
	defer closeServ()
	t := &testing.T{}
	c := client.New(endpoint)

	b.ResetTimer()
	i := 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			func() (*http.Response, *http.Response) {
				return c.UpdateSKU(t, "12", add1(i)), c.GetSKU(t, "12")
			}()

		}
	})

}

func TestPrintSKUs_StatusOK(t *testing.T) {
	printSKU := sf.GET(t, forest.Path("/products"))
	if forest.ExpectStatus(t, printSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", printSKU.StatusCode, http.StatusOK)
	}
}

func TestAddSKU_HappyPath_StatusCreated(t *testing.T) {
	uid := newUUID()
	sample := sku{ID: uid, Name: "polo", Vendor: "RL", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)

	addSKU := sf.POST(t, forest.Path("/products").Body(string(productInfo)))
	if forest.ExpectStatus(t, addSKU, http.StatusCreated) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusCreated)
	}

	path := fmt.Sprintf("/products/%s", uid)
	getSKU := sf.GET(t, forest.Path(path))
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
	uid := newUUID()
	sample := sku{ID: uid, Name: "pen", Vendor: "bic", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	addSKU := sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	if forest.ExpectStatus(t, addSKU, http.StatusCreated) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusCreated)
	}

	addSKU = sf.POST(t, forest.Path("/products").Body(string(productInfo)))
	if forest.ExpectStatus(t, addSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", addSKU.StatusCode, http.StatusOK)
	}
}

func TestDeleteID_idDoesntExist_NoContent(t *testing.T) {
	delSKU := sf.DELETE(t, forest.Path("/products/1"))
	if forest.ExpectStatus(t, delSKU, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU.StatusCode, http.StatusNoContent)
	}
}

func TestDeleteID_idExists_StatusOK(t *testing.T) {
	uid := newUUID()
	sample := sku{ID: uid, Name: "watch", Vendor: "breitling", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	path := fmt.Sprintf("/products/%s", uid)
	delSKU := sf.DELETE(t, forest.Path(path))
	if forest.ExpectStatus(t, delSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU.StatusCode, http.StatusOK)
	}

	getSKU := sf.GET(t, forest.Path(path))
	if forest.ExpectStatus(t, getSKU, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusNoContent)
	}
}

func TestDeleteID_removeSameIDTwice_StatusOK(t *testing.T) {
	uid := newUUID()
	sample := sku{ID: uid, Name: "watch", Vendor: "breitling", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	path := fmt.Sprintf("/products/%s", uid)
	delSKU := sf.DELETE(t, forest.Path(path))
	if forest.ExpectStatus(t, delSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU.StatusCode, http.StatusOK)
	}

	delSKU2 := sf.DELETE(t, forest.Path(path))
	if forest.ExpectStatus(t, delSKU2, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", delSKU2.StatusCode, http.StatusNoContent)
	}
}

func TestUpdateSKU_productExists_Success(t *testing.T) {
	uid := newUUID()
	sample := sku{ID: uid, Name: "watch", Vendor: "breitling", Quantity: 25}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	reqBody := fmt.Sprintf("{\"id\":\"%s\",\"quantity\":1000}", uid)
	updateSKU := sf.PUT(t, forest.Path("/products").Body(reqBody))
	if forest.ExpectStatus(t, updateSKU, http.StatusOK) != true {
		t.Errorf("Actual: %v\nExpected: %v", updateSKU.StatusCode, http.StatusOK)
	}

	path := fmt.Sprintf("/products/%s", uid)
	getSKU := sf.GET(t, forest.Path(path))
	body, _ := ioutil.ReadAll(getSKU.Body)
	var actual sku
	json.Unmarshal(body, &actual)

	sample.Quantity = 1000
	if actual != sample {
		t.Errorf("Actual: %v\nExpected: %v", actual, sample)
	}
}

func TestUpdateSKU_productDoesntExists_BadReqest(t *testing.T) {
	uid := newUUID()
	b := fmt.Sprintf("{\"id\":\"%s\",\"quantity\":\"1000\"}", uid)
	updateSKU := sf.PUT(t, forest.Path("/products").Body(b))
	if forest.ExpectStatus(t, updateSKU, http.StatusBadRequest) != true {
		t.Errorf("Actual: %v\nExpected: %v", updateSKU.StatusCode, http.StatusBadRequest)
	}
}

func TestGetSKU_idExists_Success(t *testing.T) {
	uid := newUUID()
	sample := sku{ID: uid, Name: "clock", Vendor: "atomic", Quantity: 120}
	productInfo, _ := json.Marshal(&sample)
	sf.POST(t, forest.Path("/products").Body(string(productInfo)))

	path := fmt.Sprintf("/products/%v", uid)
	getSKU := sf.GET(t, forest.Path(path))
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

func TestGetSKU_idDoesntExists_NoContent(t *testing.T) {
	getSKU := sf.GET(t, forest.Path("/products/404"))
	if forest.ExpectStatus(t, getSKU, http.StatusNoContent) != true {
		t.Errorf("Actual: %v\nExpected: %v", getSKU.StatusCode, http.StatusNoContent)
	}
}

func newUUID() string {
	u := uuid.NewV1()
	var i big.Int
	res, _ := i.SetString(strings.Replace(u.String(), "-", "", 4), 16)
	return res.String()
}

package benchmark

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strings"
	"testing"

	"github.com/emicklei/forest"
	"github.com/rue-brettadcock/storefront/intigration/client"
	uuid "github.com/satori/go.uuid"
)

type sku struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Vendor   string `json:"vendor,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

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

func BenchmarkRaceTest_getSKU_noRace_statusOK(b *testing.B) {
	initServ()
	defer closeServ()

	t := &testing.T{}
	c := client.New(endpoint)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp := c.GetSKU(t, "12")

			if forest.ExpectStatus(t, resp, 200) != true {
				b.Log(resp.StatusCode)
				b.Error("this broke")
			}
		}
	})
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
			i++
			go func() (*http.Response, *http.Response) {
				return c.UpdateSKU(t, "12", i), c.GetSKU(t, "12")
			}()

		}
	})

}

func newUUID() string {
	u := uuid.NewV1()
	var i big.Int
	res, _ := i.SetString(strings.Replace(u.String(), "-", "", 4), 16)
	return res.String()
}

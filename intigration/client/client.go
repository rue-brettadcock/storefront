package client

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/emicklei/forest"
)

type Invoker struct {
	url string
}

func New(endpoint string) Invoker {
	return Invoker{url: endpoint}
}

func (c Invoker) GetSKU(t *testing.T, id string) *http.Response {
	server := forest.NewClient(c.url, new(http.Client))
	path := fmt.Sprintf("/products/%v", id)
	return server.GET(t, forest.Path(path))
}

func (c Invoker) UpdateSKU(t *testing.T, id string, amt int) *http.Response {
	server := forest.NewClient(c.url, new(http.Client))
	body := fmt.Sprintf(`{"id":"%s","quantity":%s}`, id, strconv.Itoa(amt))
	return server.PUT(t, forest.Path("/products/").Body(body))
}

package service

import (
	"flag"
	"log"
	"net/http"

	"goji.io/pat"

	"github.com/rue-brettadcock/storefront/logic"
	"goji.io"
)

var (
	bindTo string
)

func init() {
	flag.StringVar(&bindTo, "listen", ":8080", "host:port to bind to")
}

// ListenAndServe initializes and starts the service
func ListenAndServe() {
	handler := Presentation{logic: logic.New()}
	mux := goji.NewMux()

	// mux.HandleFunc(pat.Get("/addSKU/:id/:name/:vendor/:quantity"), handler.addSKU)
	mux.HandleFunc(pat.Get("/addSKU/:vals"), handler.addSKU)
	mux.HandleFunc(pat.Get("/updateSKU/:id/:quantity"), handler.updateSKU)
	mux.HandleFunc(pat.Get("/printSKUs"), handler.printSKUs)
	mux.HandleFunc(pat.Get("/getSKU/:id"), handler.getSKU)
	mux.HandleFunc(pat.Get("/deleteSKU/:id"), handler.deleteSKU)

	log.Println("Listening...")
	http.ListenAndServe(bindTo, mux)
}

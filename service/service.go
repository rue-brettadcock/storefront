package service

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rue-brettadcock/storefront/logic"
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
	router := mux.NewRouter()

	router.HandleFunc("/products", handler.addSKU).Methods("POST")
	router.HandleFunc("/products/{id}", handler.getSKU).Methods("GET")
	router.HandleFunc("/products", handler.printSKUs).Methods("GET")
	router.HandleFunc("/products", handler.updateSKU).Methods("PUT")
	router.HandleFunc("/products/{id}", handler.deleteSKU).Methods("DELETE")

	log.Println("Listening...")
	http.ListenAndServe(bindTo, router)
}

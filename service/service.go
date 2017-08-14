package service

import (
	"flag"
	"net/http"

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

	http.ListenAndServe(bindTo, mux)
}

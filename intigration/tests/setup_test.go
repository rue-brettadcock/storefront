package tests

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/rue-brettadcock/storefront/service"
)

var (
	endpoint string
)

func setup() {
	flag.StringVar(&endpoint, "endpoint", "http://localhost:8080", "target endpoint")
	flag.Parse()
	fmt.Println("endpoint: ", endpoint)
	go service.ListenAndServe()
	fmt.Println("Listening ...")
}

func shutdown() {
}

// TestMain has custom setup and shutdown
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()

	os.Exit(code)
}

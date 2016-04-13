package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/context"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// version is dynamically loaded from the Makefile at compile time.
var version string

const iface = "127.0.0.1:5000"

func main() {
	// Create router & add middleware
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	// Mount handlers
	r.Get("/version", defaultHandler)
	r.Get("/", defaultHandler)
	log.Printf("Server (v%s) listening on %s\n", version, iface)
	http.ListenAndServe(iface, r)
}

// DefaultHandler returns the standard Service and Status response.
func defaultHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf(`{
		"meta": {"code": 200},
		"data": {"service": "go-to-sample", "version": "%s" }
	}`, version))
}

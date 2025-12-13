// Package cmd
package cmd

import (
	"net/http"

	"crud/handlers"
	"crud/middleware"
)

func initRoute(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
}

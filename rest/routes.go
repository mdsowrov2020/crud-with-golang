// Package cmd
package rest

import (
	"net/http"

	"crud/rest/handlers"
	"crud/rest/middleware"
)

func initRoute(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
}

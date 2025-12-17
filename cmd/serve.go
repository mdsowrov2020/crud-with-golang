// Package cmd

package cmd

import (
	"crud/config"
	"crud/rest"
	"crud/rest/handlers/product"
	"crud/rest/handlers/user"
	"crud/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	// rest.Start(cnf)
	middlewares := middleware.NewMiddlewares(cnf)
	productHandler := product.NewHandler(middlewares)
	userHandler := user.NewHandler()
	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}

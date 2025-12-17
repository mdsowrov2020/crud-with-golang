// Package cmd

package cmd

import (
	"crud/config"
	"crud/rest"
	"crud/rest/handlers/product"
	"crud/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()
	// rest.Start(cnf)
	productHandler := product.NewHandler()
	userHandler := user.NewHandler()
	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}

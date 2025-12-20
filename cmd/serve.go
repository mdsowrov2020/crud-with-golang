// Package cmd

package cmd

import (
	"crud/config"
	"crud/repo"
	"crud/rest"
	"crud/rest/handlers/product"
	"crud/rest/handlers/user"
	"crud/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	// rest.Start(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}

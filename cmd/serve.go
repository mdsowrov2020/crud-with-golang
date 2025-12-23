// Package cmd

package cmd

import (
	"fmt"
	"os"

	"crud/config"
	"crud/infra/db"
	"crud/repo"
	"crud/rest"
	"crud/rest/handlers/product"
	"crud/rest/handlers/user"
	"crud/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.GetConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// rest.Start(cnf)

	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

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

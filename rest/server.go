package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"crud/config"

	"crud/rest/handlers/product"
	"crud/rest/handlers/user"
	"crud/rest/middleware"
)

type Server struct {
	cnf            config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (server *Server) Start() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	wrappedMux := manager.WrapWithMux(mux)

	// initRoute(mux, manager)
	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoute(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HTTPPort)
	fmt.Println("Server running on port ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error startting the server", err)
	}
}

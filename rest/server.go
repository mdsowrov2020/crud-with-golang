package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"crud/config"

	"crud/rest/middleware"
)

func Server(cnf config.Config) {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	wrappedMux := manager.WrapWithMux(mux)

	initRoute(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HTTPPort)
	fmt.Println("Server running on port ", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error startting the server", err)
	}
}

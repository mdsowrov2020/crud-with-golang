// Package cmd

package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	"crud/config"
	"crud/middleware"
)

func Serve() {
	cnf := config.GetConfig()
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
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error startting the server", err)
	}
}

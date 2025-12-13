// Package cmd

package cmd

import (
	"fmt"
	"net/http"

	"crud/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)
	mux := http.NewServeMux()

	wrappedMux := manager.WrapWithMux(mux)

	initRoute(mux, manager)

	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error startting the server", err)
	}
}

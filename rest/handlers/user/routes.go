package user

import (
	"net/http"

	"crud/rest/middleware"
)

func (h *Handler) RegisterRoute(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(h.Login)))
}

// Package handlers

package product

import (
	"net/http"

	"crud/database"
	"crud/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}

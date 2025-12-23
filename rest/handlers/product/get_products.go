// Package handlers

package product

import (
	"log"
	"net/http"

	"crud/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()
	if err != nil {
		log.Printf("Error fetching products: %v", err)
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK, productList)
}

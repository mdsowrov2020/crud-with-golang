package product

import (
	"net/http"
	"strconv"

	"crud/util"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product id", 400)
		return
	}

	product, err := h.productRepo.Get(pID)
	if err != nil {

		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, http.StatusOK, product)
}

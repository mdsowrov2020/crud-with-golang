package product

import (
	"net/http"
	"strconv"

	"crud/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid product id")

		return
	}

	err = h.productRepo.Delete(pID)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	util.SendData(w, http.StatusOK, "Successfully deleted product")
}

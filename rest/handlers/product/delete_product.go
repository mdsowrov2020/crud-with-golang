package product

import (
	"net/http"
	"strconv"

	"crud/database"
	"crud/util"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product id", 400)
		return
	}

	database.Delete(pID)
	util.SendData(w, "Successfully deleted product", 200)
}

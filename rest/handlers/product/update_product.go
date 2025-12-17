package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud/database"
	"crud/util"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product id", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please provide a valid JSON", 400)
	}

	newProduct.ID = pID
	database.Update(newProduct)
	util.SendData(w, "Successfully updated product", 201)
}

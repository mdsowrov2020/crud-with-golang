package product

import (
	"encoding/json"
	"net/http"

	"crud/database"
	"crud/util"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Please provide a valid JSON", 400)
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, http.StatusCreated)
}

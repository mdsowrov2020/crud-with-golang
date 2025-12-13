package handlers

import (
	"net/http"
	"strconv"

	"crud/database"
	"crud/util"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product id", 400)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == pID {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Product not found", 404)
}

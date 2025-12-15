package handlers

import (
	"net/http"
	"strconv"

	"crud/database"
	"crud/util"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product id", 400)
		return
	}

	product := database.Get(pID)

	if product == nil {
		util.SendError(w, "Product not found", 404)
		return
	}

	util.SendData(w, product, 200)
}

package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"crud/repo"
	"crud/util"
)

type RequestUpdateProduct struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pID, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid product id")
		return
	}

	var req RequestUpdateProduct

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// req.ID = pID
	_, err = h.productRepo.Update(repo.Product{
		ID:          pID,
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	util.SendData(w, http.StatusCreated, "Successfully updated product")
}

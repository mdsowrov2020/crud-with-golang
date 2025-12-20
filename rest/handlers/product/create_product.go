package product

import (
	"encoding/json"
	"net/http"

	"crud/repo"
	"crud/util"
)

type RequestCreateProduct struct {
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageURL    string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req RequestCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Please provide a valid req body")
		return
	}

	// createdProduct := database.Store(newProduct)
	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Price:       req.Price,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error")
	}

	util.SendData(w, http.StatusCreated, createdProduct)
}

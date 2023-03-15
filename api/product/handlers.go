package product

import (
	"encoding/json"
	"net/http"

	"github.com/garoque/comprovate/database/product"
)

type ProductHandlerInterface interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	db product.Database
}

func NewProductHandler(db product.Database) ProductHandlerInterface {
	return &handler{db}
}

func (h *handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.db.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

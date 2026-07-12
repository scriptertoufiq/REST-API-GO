package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	data, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, data, 200)

}

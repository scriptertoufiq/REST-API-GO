package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	page, _ := strconv.ParseInt(reqQuery.Get("page"), 10, 64)
	limit, _ := strconv.ParseInt(reqQuery.Get("limit"), 10, 64)

	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	data, err := h.service.List(int(page), int(limit))
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	count, err := h.service.Count()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	util.SendPage(w, data, page, limit, count)
}

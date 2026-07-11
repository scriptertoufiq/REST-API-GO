package review

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetReviews(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.List(), 200)

}

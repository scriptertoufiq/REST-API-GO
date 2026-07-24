package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
	"sync"
)

var count int64

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

	var wg sync.WaitGroup
	var mg sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		mg.Lock()
		defer mg.Unlock()
		cnt, err := h.service.Count()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		count = cnt

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mg.Lock()
		defer mg.Unlock()
		count1, err := h.service.Count()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		count2, err := h.service.Count()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		println("Count1:", count1)
		println("Count2:", count2)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mg.Lock()
		defer mg.Unlock()
		count3, err := h.service.Count()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		count4, err := h.service.Count()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		println("Count3:", count3)
		println("Count4:", count4)

	}()
	//time.Sleep(8 * time.Second) // Simulate some processing time
	wg.Wait()
	util.SendPage(w, data, page, limit, count)
}

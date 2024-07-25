package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/SolBaa/go-challenge/internal/models"
	"github.com/SolBaa/go-challenge/internal/service"
)

type Handler struct {
	service service.LTPService
}

func NewHandler(s service.LTPService) Handler {
	return Handler{
		service: s,
	}

}

func (h Handler) GetLTP(w http.ResponseWriter, r *http.Request) {
	// get query params
	pair := r.URL.Query().Get("pair")
	if pair == "" {
		http.Error(w, "pair query param is required", http.StatusBadRequest)
		return
	}

	// split pairs
	pairs := strings.Split(pair, ",")
	var response models.Response
	// fetch LTP
	for _, p := range pairs {
		ltp, err := h.service.FetchLTP(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.LTP = append(response.LTP, models.LTP{Pair: p, Amount: ltp.Amount})
	}
	// fmt.Println(response)
	// return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h Handler) HealthhCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

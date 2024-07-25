package api

import (
	"github.com/SolBaa/go-challenge/internal/repository"
	"github.com/SolBaa/go-challenge/internal/service"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	repo := repository.NewKrakenRepository()
	ltpService := service.NewService(repo)
	handler := NewHandler(ltpService)
	r.Get("/api/v1/health", handler.HealthhCheck)
	r.Get("/api/v1/ltp", handler.GetLTP)
	return r
}

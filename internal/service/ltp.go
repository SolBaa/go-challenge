package service

import (
	"github.com/SolBaa/go-challenge/internal/models"
	"github.com/SolBaa/go-challenge/internal/repository"
)

type LTPService interface {
	FetchLTP(pair string) (models.LTP, error)
}

type ltpService struct {
	ltpRepo repository.LTPRepository
}

func NewService(ltpRepo repository.LTPRepository) LTPService {
	return &ltpService{
		ltpRepo: ltpRepo,
	}
}

func (s *ltpService) FetchLTP(pair string) (models.LTP, error) {
	return s.ltpRepo.FetchLTP(pair)
}

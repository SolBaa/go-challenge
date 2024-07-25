package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SolBaa/go-challenge/internal/models"
)

const krakenAPI = "https://api.kraken.com/0/public/Ticker?pair="

type LTPRepository interface {
	FetchLTP(pair string) (models.LTP, error)
}

type ltpRepository struct {
}

func NewKrakenRepository() LTPRepository {
	return &ltpRepository{}
}

func (r *ltpRepository) FetchLTP(pair string) (models.LTP, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", krakenAPI, pair))
	if err != nil {
		return models.LTP{}, err
	}
	defer resp.Body.Close()
	var krakenResp models.KrakenResponse
	if err := json.NewDecoder(resp.Body).Decode(&krakenResp); err != nil {
		return models.LTP{}, err
	}

	if len(krakenResp.Result) > 0 {
		return models.LTP{}, fmt.Errorf("error from kraken API: %v", krakenResp.Result)
	}

	ticker, exists := krakenResp.Result[pair]
	if !exists {
		return models.LTP{}, fmt.Errorf("pair %s not found", pair)
	}

	ltpValue := ticker.C[0]
	amount, err := strconv.ParseFloat(ltpValue, 64)
	if err != nil {
		return models.LTP{}, err
	}

	return models.LTP{Pair: pair, Amount: amount}, nil

}

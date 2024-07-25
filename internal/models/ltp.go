package models

type LTP struct {
	Pair   string  `json:"pair"`
	Amount float64 `json:"amount"`
}

type Response struct {
	LTP []LTP `json:"ltp"`
}

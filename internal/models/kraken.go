package models

type KrakenResponse struct {
	Result map[string]KrakenTicker `json:"result"`
}

type KrakenTicker struct {
	C []string `json:"c"`
}

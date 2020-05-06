package resources

import (
	"encoding/json"
)

const (
	DirectFxRatesPath     = "/api/v2.0.0/direct_fx/rates"
	DirectFxExchangesPath = "/api/v2.0.0/direct_fx/exchanges"
)

type DirectFxRatesResponse struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  struct {
		Ts           string      `json:"ts"`
		CurrencyPair string      `json:"currencyPair"`
		BuyRate      json.Number `json:"buyRate"`
		SellRate     json.Number `json:"sellRate"`
	} `json:"result"`
}

type DirectFxExchangesResponse struct {
	Success   bool `json:"success"`
	Duplicate bool `json:"success"`
	Errors    []Error
	Result    struct {
		ExchangeId string `json:"exchangeId"`
	} `json:"result"`
}

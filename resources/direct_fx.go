package resources

import (
	"encoding/json"
)

const DirectFxRatesPath = "/api/v2.0.0/direct_fx/rates"

type DirectFxRatesResponse struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  struct {
		Ts           string      `json:"ts"`
		CurrencyPair string      `json:"currencyPair"`
		BuyRate      json.Number `json:"buyRate"`
		SellRate     json.Number `json:"sellRate"`
	}
}

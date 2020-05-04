package resources

import (
	"encoding/json"
)

const (
	MarketFxBestOffersPath   = "/api/v2.0.0/market_fx/best_offers"
	MarketFxOrdersPath       = "/api/v2.0.0/market_fx/orders"
	MarketFxOrdersActivePath = "/api/v2.0.0/market_fx/orders/active"
)

type MarketFxBestOffers struct {
	Success bool    `json:"success"`
	Errors  []Error // Extend Error for this particular case
	Result  struct {
		Ts           string `json:"ts"`
		CurrencyPair string `json:"currencyPair"`
		Bids         []struct {
			Price                   json.Number `json:"price"`
			Volume                  json.Number `json:"volume"`
			ValueInOppositeCurrency json.Number `json:"valueInOppositeCurrency"`
		}
		Asks []struct {
			Price                   json.Number `json:"price"`
			Volume                  json.Number `json:"volume"`
			ValueInOppositeCurrency json.Number `json:"valueInOppositeCurrency"`
		}
	}
}

type MarketFxOrders struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  []struct {
		OrderId string `json:"orderId"`
		// Add remaining
	}
}

type MarketFxOrdersActive struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  []struct {
		OrderId string `json:"orderId"`
		// Add remaining
	}
}

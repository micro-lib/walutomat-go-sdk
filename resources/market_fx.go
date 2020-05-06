package resources

import (
	"encoding/json"
)

const (
	MarketFxBestOffersPath   = "/api/v2.0.0/market_fx/best_offers"
	MarketFxOrdersPath       = "/api/v2.0.0/market_fx/orders"
	MarketFxOrdersActivePath = "/api/v2.0.0/market_fx/orders/active"
	MarketFxOrdersClosePAth  = "/api/v2.0.0/market_fx/orders/close"
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
		} `json:"bids"`
		Asks []struct {
			Price                   json.Number `json:"price"`
			Volume                  json.Number `json:"volume"`
			ValueInOppositeCurrency json.Number `json:"valueInOppositeCurrency"`
		} `json:"asks"`
	} `json:"result"`
}

type MarketFxOrders struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  []struct {
		SubmitId           string `json:"submitId"`
		SubmitTs           string `json:"submitTs"`
		UpdateTs           string `json:"updateTs"`
		Status             string `json:"status"`
		Completion         int    `json:"completion"`
		CurrencyPair       string `json:"currencyPair"`
		BuySell            string `json:"buySell"`
		Volume             string `json:"volume"`
		VolumeCurrency     string `json:"volumeCurrency"`
		LimitPrice         string `json:"limitPrice"`
		SoldAmount         string `json:"soldAmount"`
		SoldCurrency       string `json:"soldCurrency"`
		BoughtAmount       string `json:"boughtAmount"`
		BoughtCurrency     string `json:"boughtCurrency"`
		CommissionAmount   string `json:"commissionAmount"`
		CommissionCurrency string `json:"commissionCurrency"`
		CommissionRate     string `json:"commissionRate"`
	} `json:"result"`
}

type MarketFxOrdersActive struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  []struct {
		OrderId            string `json:"orderId"`
		SubmitId           string `json:"submitId"`
		SubmitTs           string `json:"submitTs"`
		UpdateTs           string `json:"updateTs"`
		Status             string `json:"status"`
		Completion         int    `json:"completion"`
		CurrencyPair       string `json:"currencyPair"`
		BuySell            string `json:"buySell"`
		Volume             string `json:"volume"`
		VolumeCurrency     string `json:"volumeCurrency"`
		LimitPrice         string `json:"limitPrice"`
		SoldAmount         string `json:"soldAmount"`
		SoldCurrency       string `json:"soldCurrency"`
		BoughtAmount       string `json:"boughtAmount"`
		BoughtCurrency     string `json:"boughtCurrency"`
		CommissionAmount   string `json:"commissionAmount"`
		CommissionCurrency string `json:"commissionCurrency"`
		CommissionRate     string `json:"commissionRate"`
	}
}

type MarketFxOrdersResponse struct {
	Success   bool `json:"success"`
	Errors    []Error
	Duplicate bool `json:"duplicate"`
	Result    struct {
		OrderId string `json:"orderId"`
	} `json:"result"`
}

type MarketFxOrdersCloseResponse struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  struct {
		OrderId            string `json:"orderId"`
		SubmitId           string `json:"submitId"`
		SubmitTs           string `json:"submitTs"`
		UpdateTs           string `json:"updateTs"`
		Status             string `json:"status"`
		Completion         int    `json:"completion"`
		CurrencyPair       string `json:"currencyPair"`
		BuySell            string `json:"buySell"`
		Volume             string `json:"volume"`
		VolumeCurrency     string `json:"volumeCurrency"`
		LimitPrice         string `json:"limitPrice"`
		SoldAmount         string `json:"soldAmount"`
		SoldCurrency       string `json:"soldCurrency"`
		BoughtAmount       string `json:"boughtAmount"`
		BoughtCurrency     string `json:"boughtCurrency"`
		CommissionAmount   string `json:"commissionAmount"`
		CommissionCurrency string `json:"commissionCurrency"`
		CommissionRate     string `json:"commissionRate"`
	} `json:"result"`
}

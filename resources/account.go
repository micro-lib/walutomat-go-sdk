package resources

import (
	"encoding/json"
)

const (
	AccountBalancesPath = "/api/v2.0.0/account/balances"
	AccountHistoryPath  = "/api/v2.0.0/account/history"
)

type AccountBalancesResponse struct {
	Success bool `json:"success"`
	Errors  []Error
	Result  []struct {
		Currency         string      `json:"currency"`
		BalanceTotal     json.Number `json:"balanceTotal"`
		BalanceAvailable json.Number `json:"balanceAvailable"`
		BalanceReserved  json.Number `json:"balanceReserved"`
	}
}

type AccountHistoryResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		HistoryItemId json.Number `json:"historyItemId"`
		// Add remaining fields
	}
}

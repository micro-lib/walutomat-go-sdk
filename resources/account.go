package resources

import (
	"encoding/json"
	"github.com/micro-lib/walutomat-go-sdk"
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
	} `json:"result"`
}

type AccountHistoryResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		HistoryItemId    int    `json:"historyItemId"`
		TransactionId    string `json:"transactionId"`
		Ts               string `json:"ts"`
		OperationAmount  string `json:"operationAmount"`
		BalanceAfter     string `json:"balanceAfter"`
		OperationDetails []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		Currency              string                  `json:"currency"`
		OperationType         walutomat.OperationType `json:"operationType"`
		OperationDetailedType string                  `json:"operationDetailedType"`
		SubmitId              string                  `json:"submitId"`
	} `json:"result"`
}

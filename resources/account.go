package resources

import (
	"encoding/json"
	"github.com/micro-lib/walutomat-go-sdk/types"
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
		Currency              string              `json:"currency"`
		OperationType         types.OperationType `json:"operationType"`
		OperationDetailedType string              `json:"operationDetailedType"`
		SubmitId              string              `json:"submitId"`
	} `json:"result"`
}

type SortOrder string

const (
	SortAsc  SortOrder = "ASC"
	SortDesc SortOrder = "DESC"
)

type AccountHistoryQuery struct {
	Currencies    []types.Currency
	OperationType types.OperationType
	ItemLimit     int
	ContinueFrom  int
	SortOrder     SortOrder
	DateFrom      string
	DateTo        string
}

func NewAccountHistoryQueryWithDefaults() AccountHistoryQuery {
	return AccountHistoryQuery{
		Currencies:    nil,
		OperationType: "",
		ItemLimit:     200,
		ContinueFrom:  0,
		SortOrder:     SortAsc,
		DateFrom:      "",
		DateTo:        "",
	}
}

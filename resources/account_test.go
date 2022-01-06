package resources

import (
	"encoding/json"
	"github.com/micro-lib/walutomat-go-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccountHistory(t *testing.T) {
	var rawResponse = `{
  "success": true,
  "result": [
    {
      "historyItemId": 51277549,
      "transactionId": "a4570b2b-23cf-4703-baf1-9b6361be267c",
      "ts": "2018-10-08T13:10:14.788Z",
      "operationAmount": "-432.43 PLN",
      "balanceAfter": "-423.75 PLN",
      "operationDetails": [
        {
          "key": "amountSold",
          "value": "432.43 PLN"
        },
        {
          "key": "amountBought",
          "value": "100.00 EUR"
        },
        {
          "key": "rate",
          "value": "4.3243"
        },
        {
          "key": "rateTs",
          "value": "2018-10-08T13:10:00.000Z"
        }
      ],
      "currency": "PLN",
      "operationType": "DIRECT_FX",
      "operationDetailedType": "DIRECT_FX",
      "submitId": "test105"
    },
    {
      "historyItemId": 51277650,
      "transactionId": "8981849f-30ea-4049-9ecc-1a25db0920ae",
      "ts": "2018-10-08T13:11:41.970Z",
      "operationAmount": "-0.24 EUR",
      "balanceAfter": "105.09 EUR",
      "operationDetails": [
        {
          "key": "amountSold",
          "value": "0.24 EUR"
        },
        {
          "key": "amountBought",
          "value": "1.00 PLN"
        },
        {
          "key": "rate",
          "value": "4.2943"
        },
        {
          "key": "rateTs",
          "value": "2018-10-08T13:10:00.000Z"
        }
      ],
      "currency": "EUR",
      "operationType": "DIRECT_FX",
      "operationDetailedType": "DIRECT_FX",
      "submitId": "test106"
    }
  ]
}`

	var response = AccountHistoryResponse{}
	err := json.Unmarshal([]byte(rawResponse), &response)

	assert.Nil(t, err)

	var expectedResponse = AccountHistoryResponse{
		Success: true,
		Result: []struct {
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
		}{{
			HistoryItemId:         51277549,
			TransactionId:         "a4570b2b-23cf-4703-baf1-9b6361be267c",
			Ts:                    "2018-10-08T13:10:14.788Z",
			OperationAmount:       "-432.43 PLN",
			BalanceAfter:          "-423.75 PLN",
			Currency:              "PLN",
			OperationType:         types.DirectFX,
			OperationDetailedType: "DIRECT_FX",
			OperationDetails: []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{{
				Key:   "amountSold",
				Value: "432.43 PLN",
			}, {
				Key:   "amountBought",
				Value: "100.00 EUR",
			}, {
				Key:   "rate",
				Value: "4.3243",
			}, {
				Key:   "rateTs",
				Value: "2018-10-08T13:10:00.000Z",
			}},
			SubmitId: "test105",
		}, {
			HistoryItemId:         51277650,
			TransactionId:         "8981849f-30ea-4049-9ecc-1a25db0920ae",
			Ts:                    "2018-10-08T13:11:41.970Z",
			OperationAmount:       "-0.24 EUR",
			BalanceAfter:          "105.09 EUR",
			Currency:              "EUR",
			OperationType:         types.DirectFX,
			OperationDetailedType: "DIRECT_FX",
			OperationDetails: []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{{
				Key:   "amountSold",
				Value: "0.24 EUR",
			}, {
				Key:   "amountBought",
				Value: "1.00 PLN",
			}, {
				Key:   "rate",
				Value: "4.2943",
			}, {
				Key:   "rateTs",
				Value: "2018-10-08T13:10:00.000Z",
			}},
			SubmitId: "test106",
		}},
	}
	assert.Exactly(t, &expectedResponse, &response)
}

package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarketFxOrdersActive(t *testing.T) {
	var rawResponse = `{
  "success": true,
  "errors": [
    {
      "key": "INSUFFICIENT_FUNDS",
      "description": "Requested payout exceeds current wallet balance {walletBalance}",
      "errorData": [
        {
          "key": "walletBalance",
          "value": "3.14 EUR"
        }
      ]
    }
  ],
  "result": [
    {
      "orderId": "2035e361-e672-457a-9c3c-0e86e5ff54d6",
      "submitId": "916f1f98-01f6-412a-85e7-2482f1f4c112",
      "submitTs": "2018-02-02T10:06:01.111Z",
      "updateTs": "2018-02-02T10:06:01.396Z",
      "status": "ACTIVE",
      "completion": 67,
      "currencyPair": "EURPLN",
      "buySell": "BUY",
      "volume": "999.00",
      "volumeCurrency": "EUR",
      "limitPrice": "4.2321",
      "soldAmount": "999.00",
      "soldCurrency": "EUR",
      "boughtAmount": "999.00",
      "boughtCurrency": "EUR",
      "commissionAmount": "9.00",
      "commissionCurrency": "EUR",
      "commissionRate": "0.0020"
    }
  ]
}`
	var response = MarketFxOrdersActive{}
	json.Unmarshal([]byte(rawResponse), &response)

	var expectedResponse = MarketFxOrdersActive{
		Success: true,
		Errors: []Error{Error{
			Key:         "INSUFFICIENT_FUNDS",
			Description: "Requested payout exceeds current wallet balance {walletBalance}",
			ErrorData: []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{
				{
					Key:   "walletBalance",
					Value: "3.14 EUR",
				},
			},
		}},
		Result: []struct {
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
		}{{
			OrderId:            "2035e361-e672-457a-9c3c-0e86e5ff54d6",
			SubmitId:           "916f1f98-01f6-412a-85e7-2482f1f4c112",
			SubmitTs:           "2018-02-02T10:06:01.111Z",
			UpdateTs:           "2018-02-02T10:06:01.396Z",
			Status:             "ACTIVE",
			Completion:         67,
			CurrencyPair:       "EURPLN",
			BuySell:            "BUY",
			Volume:             "999.00",
			VolumeCurrency:     "EUR",
			LimitPrice:         "4.2321",
			SoldAmount:         "999.00",
			SoldCurrency:       "EUR",
			BoughtAmount:       "999.00",
			BoughtCurrency:     "EUR",
			CommissionAmount:   "9.00",
			CommissionCurrency: "EUR",
			CommissionRate:     "0.0020",
		}},
	}

	assert.Exactly(t, &expectedResponse, &response)
}

func TestMarketFxOrdersCloseResponseDecode(t *testing.T) {

	var rawResponse = `{
  "success": true,
  "errors": [
    {
      "key": "INSUFFICIENT_FUNDS",
      "description": "Requested payout exceeds current wallet balance {walletBalance}",
      "errorData": [
        {
          "key": "walletBalance",
          "value": "3.14 EUR"
        }
      ]
    }
  ],
  "result": {
    "orderId": "2035e361-e672-457a-9c3c-0e86e5ff54d6",
    "submitId": "916f1f98-01f6-412a-85e7-2482f1f4c112",
    "submitTs": "2018-02-02T10:06:01.111Z",
    "updateTs": "2018-02-02T10:06:01.396Z",
    "status": "ACTIVE",
    "completion": 67,
    "currencyPair": "EURPLN",
    "buySell": "BUY",
    "volume": "999.00",
    "volumeCurrency": "EUR",
    "limitPrice": "4.2321",
    "soldAmount": "999.00",
    "soldCurrency": "EUR",
    "boughtAmount": "999.00",
    "boughtCurrency": "EUR",
    "commissionAmount": "9.00",
    "commissionCurrency": "EUR",
    "commissionRate": "0.0020"
  }
}`

	var response = MarketFxOrdersCloseResponse{}

	json.Unmarshal([]byte(rawResponse), &response)

	var expectedResponse = MarketFxOrdersCloseResponse{
		Success: true,
		Errors: []Error{Error{
			Key:         "INSUFFICIENT_FUNDS",
			Description: "Requested payout exceeds current wallet balance {walletBalance}",
			ErrorData: []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			}{
				{
					Key:   "walletBalance",
					Value: "3.14 EUR",
				},
			},
		}},
		Result: struct {
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
		}{
			OrderId:            "2035e361-e672-457a-9c3c-0e86e5ff54d6",
			SubmitId:           "916f1f98-01f6-412a-85e7-2482f1f4c112",
			SubmitTs:           "2018-02-02T10:06:01.111Z",
			UpdateTs:           "2018-02-02T10:06:01.396Z",
			Status:             "ACTIVE",
			Completion:         67,
			CurrencyPair:       "EURPLN",
			BuySell:            "BUY",
			Volume:             "999.00",
			VolumeCurrency:     "EUR",
			LimitPrice:         "4.2321",
			SoldAmount:         "999.00",
			SoldCurrency:       "EUR",
			BoughtAmount:       "999.00",
			BoughtCurrency:     "EUR",
			CommissionAmount:   "9.00",
			CommissionCurrency: "EUR",
			CommissionRate:     "0.0020",
		},
	}

	assert.Exactly(t, expectedResponse, response)
}

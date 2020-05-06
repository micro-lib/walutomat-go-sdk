package resources

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIbanResponseDecode(t *testing.T) {
	var rawResponse = `{
	"success": true,
		"duplicate": false,
		"result": {
			"feeAmount": "0.23",
			"feeCurrency": "EUR",
			"transferId": "1f38762e-48e7-447d-afea-85211acb7d0e"
	}
}`
	var response = TransferIbanResponse{}

	json.Unmarshal([]byte(rawResponse), &response)

	var expectedResponse = TransferIbanResponse{
		Success:   true,
		Duplicate: false,
		Result: struct {
			FeeAmount   string `json:"feeAmount"`
			FeeCurrency string `json:"feeCurrency"`
			TransferId  string `json:"transferId"`
		}{
			FeeAmount:   "0.23",
			FeeCurrency: "EUR",
			TransferId:  "1f38762e-48e7-447d-afea-85211acb7d0e",
		},
	}

	assert.Exactly(t, response, expectedResponse)
}

func TestInternalResponseDecode(t *testing.T) {
	var rawResponse = `{
	"success": true,
		"duplicate": true,
		"result": {
		"feeAmount": "0.24",
			"feeCurrency": "EUR",
			"transferId": "1f38762e-48e7-447d-afea-85211acb7dqe"
	}
}`
	var expectedResponse = &TransferInternalResponse{
		Success:   true,
		Duplicate: true,
		Result: struct {
			FeeAmount   string `json:"feeAmount"`
			FeeCurrency string `json:"feeCurrency"`
			TransferId  string `json:"transferId"`
		}{
			FeeAmount:   "0.24",
			FeeCurrency: "EUR",
			TransferId:  "1f38762e-48e7-447d-afea-85211acb7dqe",
		},
	}

	var response = &TransferInternalResponse{}

	json.Unmarshal([]byte(rawResponse), response)

	assert.Exactly(t, expectedResponse, response)
}

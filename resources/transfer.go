package resources

const (
	TransferInternalPath = "/api/v2.0.0/transfer/internal"
	TransferIbanPath     = "/api/v2.0.0/transfer/iban"
)

type TransferInternalResponse struct {
	Success   bool `json:"success"`
	Duplicate bool `json:"duplicate"`
	Result    struct {
		FeeAmount   string `json:"feeAmount"`
		FeeCurrency string `json:"feeCurrency"`
		TransferId  string `json:"transferId"`
	} `json:"result"`
}

type TransferIbanResponse struct {
	Success   bool `json:"success"`
	Duplicate bool `json:"duplicate"`
	Result    struct {
		FeeAmount   string `json:"feeAmount"`
		FeeCurrency string `json:"feeCurrency"`
		TransferId  string `json:"transferId"`
	} `json:"result"`
}

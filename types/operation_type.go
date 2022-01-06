package types

type OperationType string

const (
	PayIn      OperationType = "PAYIN"
	PayOut     OperationType = "PAYOUT"
	Commission OperationType = "COMMISSION"
	DirectFX   OperationType = "DIRECT_FX"
	MarketFX   OperationType = "MARKET_FX"
	Transfer   OperationType = "TRANSFER"
	Other      OperationType = "OTHER"
)

package models

type OrderRequest struct {
	OrderID string `json:"orderId"`
}

type ValidateOrderResult struct {
	OrderID string
	ItemID  string
	Price   int64
}

type ReserveInventoryResult struct {
	OrderID       string
	ReservationID string
	ItemID        string
	StockDeducted bool
}

type ChargePaymentResult struct {
	OrderID  string
	ChargeID string
	Amount   int64
	Status   string
}

type SendEmailResult struct {
	OrderID  string
	ChargeID string
	Email    string
	Sent     bool
}

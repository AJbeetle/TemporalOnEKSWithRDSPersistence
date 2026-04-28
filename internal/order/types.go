package order

type OrderRequest struct {
	OrderID string `json:"orderId"`
}

type OrderResult struct {
	OrderID string
	Status  string
}

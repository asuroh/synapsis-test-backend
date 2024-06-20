package request

type TransactionDetailRequest struct {
	TransactionID string `json:"transaction_id"`
	ProductID     string `json:"product_id"`
	Qty           int64  `json:"qty"`
}

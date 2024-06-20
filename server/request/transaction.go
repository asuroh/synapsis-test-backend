package request

// TransactionRequest ...
type TransactionRequest struct {
	UserID     string   `json:"user_id"`
	UserCartID []string `json:"user_cart_id"`
}

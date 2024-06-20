package viewmodel

// TransactionVM ...
type TransactionVM struct {
	ID                string                `json:"id"`
	UserID            string                `json:"user_id"`
	Total             float64               `json:"total"`
	Status            string                `json:"status"`
	Code              string                `json:"code"`
	TransactionDetail []TransactionDetailVM `json:"transaction_detail"`
	CreatedAt         string                `json:"created_at"`
	UpdatedAt         string                `json:"updated_at"`
	DeletedAt         string                `json:"deleted_at"`
}

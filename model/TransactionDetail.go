package model

import (
	"database/sql"
	"synapsis-test-backend/usecase/viewmodel"
	"time"
)

var (
	// DefaultTransactioDetailnBy ...
	DefaultTransactioDetailnBy = "def.updated_at"
	// TransactionDetailBy ...
	TransactionDetailBy = []string{
		"def.created_at", "def.updated_at",
	}
)

// transactionDetailModel ...
type transactionDetailModel struct {
	DB *sql.DB
}

// ItransactionDetail ...
type ItransactionDetail interface {
	Store(id string, body viewmodel.TransactionDetailVM, changedAt time.Time) error
}

// TransactionDetailEntity ....
type TransactionDetailEntity struct {
	ID            string         `db:"id"`
	TransactionID string         `db:"transaction_id"`
	ProductName   string         `db:"product_name"`
	ProductID     string         `db:"product_id"`
	Price         float64        `db:"price"`
	Qty           int64          `db:"qty"`
	CreatedAt     string         `db:"created_at"`
	UpdatedAt     string         `db:"updated_at"`
	DeletedAt     sql.NullString `db:"deleted_at"`
}

// NewTransactionDetailModel ...
func NewTransactionDetailModel(db *sql.DB) ItransactionDetail {
	return &transactionDetailModel{DB: db}
}

// Store ...
func (model transactionDetailModel) Store(id string, body viewmodel.TransactionDetailVM, changedAt time.Time) (err error) {
	sql := `INSERT INTO transaction_details (id, transaction_id, product_name, product_id, price, qty, created_at, updated_at
		) VALUES(?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = model.DB.Exec(sql, id, body.TransactionID, body.ProductName, body.ProductID, body.Price, body.Qty, changedAt, changedAt)

	return err
}

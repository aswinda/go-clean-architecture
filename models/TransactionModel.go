package models

type TransactionModel struct {
	ID         int    `json:"id" validate:"numeric"`
	CustomerID int    `json:"customer_id" validate:"required,numeric"`
	EventID    int    `json:"event_id" validate:"required,numeric"`
	TotalPrice int    `json:"total_price" validate:"required,numeric"`
	Status     int    `json:"status" validate:"numeric"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

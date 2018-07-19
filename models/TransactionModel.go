package models

type TransactionModel struct {
	ID          int    `json:"id" validate:"numeric"`
	CustomerID  int    `json:"customer_id" validate:"required,numeric"`
	EventID     int    `json:"event_id" validate:"required,numeric"`
	TotalAmount int    `json:"total_amount" validate:"required,numeric"`
	TotalPrice  int    `json:"total_price" validate:"required,numeric"`
	Status      int    `json:"status" validate:"numeric"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TransactionDetailModel struct {
	ID            int    `json:"id" validate:"numeric"`
	TransactionID int    `json:"transaction_id" validate:"required,numeric"`
	TicketID      int    `json:"ticket_id" validate:"required,numeric"`
	Amount        int    `json:"amount" validate:"required,numeric"`
	TotalPrice    int    `json:"total_price" validate:"required,numeric"`
	Status        int    `json:"status" validate:"numeric"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TransactionPurchase struct {
	CustomerID int   `json:"customer_id" validate:"required,numeric"`
	EventID    int   `json:"event_id" validate:"required,numeric"`
	TicketID   []int `json:"ticket_id" validate:"required"`
	Amount     []int `json:"amount" validate:"required"`
}

type TransactionDetailViewModel struct {
	ID            int             `json:"id" validate:"numeric"`
	TransactionID int             `json:"customer_id" validate:"required,numeric"`
	Ticket        TicketViewModel `json:"ticket" validate:"required,numeric"`
	Amount        int             `json:"amount" validate:"required,numeric"`
	TotalPrice    int             `json:"total_price" validate:"required,numeric"`
}

type TransactionInfoViewModel struct {
	ID          int                           `json:"id" validate:"numeric"`
	CustomerID  int                           `json:"customer_id" validate:"required"`
	EventName   string                        `json:"event_name" validate:"required"`
	TotalAmount int                           `json:"total_amount" validate:"required,numeric"`
	TotalPrice  int                           `json:"total_price" validate:"required,numeric"`
	Details     []*TransactionDetailViewModel `json:"details"`
}

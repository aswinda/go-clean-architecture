package models

type TicketModel struct {
	ID           int    `json:"id" validate:"numeric"`
	EventID      int    `json:"event_id" validate:"required,numeric"`
	TicketTypeID int    `json:"ticket_type_id" validate:"required,numeric"`
	Quota        int    `json:"quota" validate:"required,numeric"`
	Price        int    `json:"price" validate:"required,numeric"`
	Status       int    `json:"status" validate:"numeric"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TicketViewModel struct {
	ID           int `json:"id" validate:"numeric"`
	TicketTypeID int `json:"ticket_type_id" validate:"required,numeric"`
	Quota        int `json:"quota" validate:"required,numeric"`
	Price        int `json:"price" validate:"required,numeric"`
}

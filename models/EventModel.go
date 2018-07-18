package models

type EventModel struct {
	ID          int    `json:"id" validate:"numeric"`
	Name        string `json:"name" validate:"required,gt=2,lte=45"`
	Description string `json:"description" validate:"required,gt=2,lte=500"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
	Status      int    `json:"status" validate:"numeric"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

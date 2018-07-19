package models

type LocationModel struct {
	ID        int    `json:"id" validate:"numeric"`
	Name      string `json:"name" validate:"required,gt=2,lte=45"`
	Status    int    `json:"status" validate:"numeric"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

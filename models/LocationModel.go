package models

type LocationModel struct {
	Id         int    `json:"id" validate:"numeric"`
	Name       string `json:"name" validate:"required,gt=2,lte=45"`
	Status     int    `json:"status" validate:"numeric"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

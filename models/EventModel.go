package models

type EventModel struct {
	Id          int    `json:"id" validate:"numeric"`
	Name        string `json:"name" validate:"required,gt=2,lte=45"`
	Description string `json:"description" validate:"required,gt=2,lte=500"`
	Start_time  string `json:"start_time" validate:"required"`
	End_time    string `json:"end_time" validate:"required"`
	Status      int    `json:"status" validate:"numeric"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

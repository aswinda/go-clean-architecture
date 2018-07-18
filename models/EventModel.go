package models

type EventModel struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Start_time  string `json:"start_time"`
	End_time    string `json:"end_time"`
	Status      int    `json:"status"`
	Created_at  string `json:"created_at"`
	Updated_at  string `json:"updated_at"`
}

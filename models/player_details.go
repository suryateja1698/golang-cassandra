package models

import "time"

type PlayerDetails struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Age         int       `json:"age"`
	Country     string    `json:"country"`
	Position    string    `json:"position"`
	JoinedYear  string    `json:"joined_year"`
	JoinedMonth string    `json:"joined_month"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

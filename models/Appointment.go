package models

type Appointment struct {
	Ended_at string `json:"ended_at"`
	ID     int  `json:"id"`
	User_id int `json:"user_id"`
	Started_at string `json:"started_at"`
	Trainer_id int `json:"trainer_id"`
}

type UserID struct {
	id int
}
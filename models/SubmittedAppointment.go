package models

type SubmittedAppointment struct {
	Ends_at string `json:"ends_at" validate:"required"`
	User_id int `json:"user_id" validate:"required"`
	Starts_at string `json:"starts_at" validate:"required"`
	Trainer_id int `json:"trainer_id" validate:"required"`
}
package models

// number of attributes may be expanded in the future, but for not this is 
// primarily for tracking/validating user ids

type User struct {
	User_id int `json:"user_id"`
}
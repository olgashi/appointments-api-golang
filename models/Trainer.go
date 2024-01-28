package models

// number of attributes may be expanded in the future, but for now this is 
// primarily for tracking/vaildating trainer ids

type Trainer struct {
	Trainer_id int `json:"trainer_id"`
}
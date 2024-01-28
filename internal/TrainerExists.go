package internal

import (
	"appointments-api/models"
)

func TrainerExists(id int, trainers []models.Trainer) bool {
	for _, trainer := range trainers {
		if trainer.Trainer_id == id {
			return true
		}
	}
	return false
}
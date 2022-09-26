package internal

import (
	"testing"
	"appointments-api/models"
)

var trainers = []models.Trainer{
	{
		Trainer_id: 10,
	},
	{
		Trainer_id: 20,
	},
	{
		Trainer_id: 30,
	},
}

func TestTrainerExists(t *testing.T) {
	trainerId := 10
	trainerExistsInList := TrainerExists(trainerId, trainers)

	if trainerExistsInList != true {
		t.Fatalf(`Should return true for trainer id %v, but returned %v`, trainerId, trainerExistsInList)
	}
}

func TestTrainerExists_TrainerDoesntExist(t *testing.T) {
	trainerId := 40
	trainerExistsInList := TrainerExists(trainerId, trainers)

	if trainerExistsInList != false {
		t.Fatalf(`Should return false for trainer id %v, but returned %v`, trainerId, trainerExistsInList)
	}
}
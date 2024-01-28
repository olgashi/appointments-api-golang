package internal

import (
	"testing"
	"appointments-api/models"
)

var appointments = []models.Appointment{
	{
		Ended_at: "2019-01-24T09:30:00-08:00",
				ID: 100,
				User_id: 1,
				Started_at: "2019-01-24T09:00:00-08:00",
				Trainer_id: 2,
		},
		{
				Ended_at: "2019-01-24T10:30:00-08:00",
				ID: 101,
				User_id: 2,
				Started_at: "2019-01-24T10:00:00-08:00",
				Trainer_id: 1,
		},
	}

func TestGetLargestExistingId(t *testing.T) {
	largestId := GetLargestExistingId(appointments)

	if largestId != 101 {
		t.Fatalf(`Appointment id should be equal to %v, found %v`, "101", largestId)
	}
}

func TestGenerateAppointmentId(t *testing.T) {
	nextId := GenerateAppointmentId(appointments)

	if nextId != 102 {
		t.Fatalf(`Appointment id should be equal to %v, found %v`, "102", nextId)
	}
}
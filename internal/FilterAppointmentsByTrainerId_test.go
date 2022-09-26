package internal

import (
    "testing"
		"appointments-api/models"
)

func TestFilterAppointmentsByTrainerId(t *testing.T) {
	var appointments = []models.Appointment{
		{
			Ended_at: "2019-01-24T09:30:00-08:00",
					ID: 123456789,
					User_id: 1,
					Started_at: "2019-01-24T09:00:00-08:00",
					Trainer_id: 2,
			},
			{
					Ended_at: "2019-01-24T10:30:00-08:00",
					ID: 987654321,
					User_id: 2,
					Started_at: "2019-01-24T10:00:00-08:00",
					Trainer_id: 1,
			},
		}
  filteredAppointment := FilterAppointmentsByTrainerId("1", appointments)
	apptId := filteredAppointment[0].ID

	if apptId != 987654321 {
		t.Fatalf(`Appointment id should be equal to %v, found %v`, "987654321", apptId)
	}
}
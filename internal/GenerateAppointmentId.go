package internal

import (
	"appointments-api/models"
)

func GetLargestExistingId(appointments []models.Appointment) int {
	maxId := -1
	for _, appt := range appointments {
		if appt.ID > maxId {
			maxId = appt.ID
		}
	}
	return maxId
}

func GenerateAppointmentId(appointments []models.Appointment) int {
	return GetLargestExistingId(appointments) + 1
}
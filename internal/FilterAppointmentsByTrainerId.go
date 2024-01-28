package internal

import (
	"appointments-api/models"
	"strconv"
)

func FilterAppointmentsByTrainerId(id string, appointments[] models.Appointment) []models.Appointment {
	matchingAppointments := []models.Appointment{}
	trainerIdStr, _ := strconv.Atoi(id)
  for _, app := range appointments {
		if app.Trainer_id == trainerIdStr {
			matchingAppointments = append(matchingAppointments, app)
		}
  }
	return matchingAppointments
}
package internal

import (
	models "appointments-api/models"
)

func CreateNewAppointment(endedAt string, id int, userId int, startedAt string, trainerId int) *models.Appointment {
	app := models.Appointment{Ended_at: endedAt, ID: id, User_id: userId, Started_at: startedAt, Trainer_id: trainerId}
  return &app
}
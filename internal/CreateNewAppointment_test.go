package internal

import (
    "testing"
)

func TestCreateNewAppointment(t *testing.T) {
  newAppointment := CreateNewAppointment("2019-01-24T09:30:00-08:00", 1, 1, "2019-01-24T09:00:00-08:00", 1) 

	if newAppointment.Ended_at != "2019-01-24T09:30:00-08:00" {
		t.Fatalf(`newAppointment.Ended_at should be equal to %v, found %v`, "2019-01-24T09:30:00-08:00", newAppointment.Ended_at)
	}
	if newAppointment.Started_at != "2019-01-24T09:00:00-08:00" {
		t.Fatalf(`newAppointment.Started_at should be equal to %v, found %v`, "2019-01-24T09:00:00-08:00", newAppointment.Started_at)
	}
	if newAppointment.User_id != 1 {
		t.Fatalf(`newAppointment.User_id should be equal to %v, found %v`, "2019-01-24T09:30:00-08:00", newAppointment.User_id)
	}
	if newAppointment.Trainer_id != 1 {
		t.Fatalf(`newAppointment.Trainer_id should be equal to %v, found %v`, "2019-01-24T09:30:00-08:00", newAppointment.Trainer_id)
	}
}
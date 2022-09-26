package internal

import (
	"testing"
)

func TestIsAlreadyBooked_AppointmentIsBooked(t *testing.T) {
	appointmentIsAlreadyBooked := IsAlreadyBooked("2019-01-24T09:00:00-08:00", "2019-01-24T09:30:00-08:00", "1")
	if appointmentIsAlreadyBooked != true {
		t.Fatalf(`Should return true for the timeslot that is booked but returned %v`, appointmentIsAlreadyBooked)
	}
}

func TestIsAlreadyBooked_AppointmentIsNotBooked(t *testing.T) {
	appointmentIsAlreadyBooked := IsAlreadyBooked("2022-01-24T09:00:00-08:00", "2022-01-24T09:30:00-08:00", "1")
	if appointmentIsAlreadyBooked != false {
		t.Fatalf(`Should return false for the timeslot that is booked but returned %v`, appointmentIsAlreadyBooked)
	}
}
package internal

import (
	"time"
	"appointments-api/mocks"
)

func IsAlreadyBooked(startT string, endT string, trainerId string) bool {
	startTParsed, _ := time.Parse(time.RFC3339, startT)
	endTParsed, _ := time.Parse(time.RFC3339, endT)
	appointmentsForThisTrainer := FilterAppointmentsByTrainerId(trainerId, mocks.Appointments)
	
  for _, bookedAppt := range appointmentsForThisTrainer {
		apptStart, _ := time.Parse(time.RFC3339, bookedAppt.Started_at)
		apptEnd, _ := time.Parse(time.RFC3339, bookedAppt.Ended_at)

		if (apptStart.Equal(startTParsed) && apptEnd.Equal(endTParsed)) {
			return true
		}
	}

	return false
}
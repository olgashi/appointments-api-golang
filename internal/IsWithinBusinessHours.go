package internal

import (
	"time"
	"fmt"
)

func IsWithinBusinessHours(startAppointmentStr string, endAppointmentStr string) bool {
	const startBusinessHour = 8
	const endBusinessHour = 17
	const pasicificTimezone = "America/Los_Angeles"

	loc, _ := time.LoadLocation(pasicificTimezone)
  startTime, err1 := time.Parse(time.RFC3339, startAppointmentStr)
  endTime, err2 := time.Parse(time.RFC3339, endAppointmentStr)

	if err1 != nil || err2 != nil {
		fmt.Println("Could not parse time:", err1, err2)
	}

	if (startTime.In(loc).Hour() >= startBusinessHour && endTime.In(loc).Hour() < endBusinessHour) {
		return true
	}

	return false
}
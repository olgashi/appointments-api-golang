package internal

import (
	"fmt"
	"strings"
	"time"
	"appointments-api/mocks"
)

func GetAllAvailableAppointmentsForDateRange(startsAtDate string, endsAtDate string, trainerId string) []map[string]string {

	bookedApptWithinGivenDateRange := map[string]string{}

	rangeStart, err1 := time.Parse(time.RFC3339, startsAtDate)
	rangeEnd, err2 := time.Parse(time.RFC3339, endsAtDate)

	if err1 != nil || err2 != nil {
		fmt.Errorf("Could not parse time. Errors %v, %v", err1, err2)
	}

	bookedApptList := FilterAppointmentsByTrainerId(trainerId, mocks.Appointments)

	for _, appt := range bookedApptList {
		apptStart, _ := time.Parse(time.RFC3339, appt.Started_at)
		apptEnd, _ := time.Parse(time.RFC3339, appt.Ended_at)

		if ((apptStart.After(rangeStart) || apptStart.Equal(rangeStart)) && (apptEnd.Before(rangeEnd) || apptEnd.Equal(rangeEnd))) {
			bookedApptWithinGivenDateRange[appt.Started_at] = appt.Ended_at
		}
	}

	openAppointmentSlots := []map[string]string{}

	allPossibleTimeSlots := []string{"08:00", "08:30", "09:00", "09:30", "10:00", "10:30", "11:00", "11:30",
		"12:00", "12:30", "13:00", "13:30", "14:00", "14:30", "15:00", "15:30", "16:00", "16:30", "17:00",
	}
	
	startTimeIndex := 0
	endTimeIndex := 1

	extractedStartDate := strings.Split(startsAtDate, "T")[0]
	extractedEndsAtDate := strings.Split(endsAtDate, "T")[0]
	
	for endTimeIndex < len(allPossibleTimeSlots) {
		newSlot := map[string]string{}
    currentSlotStart := extractedStartDate + "T" + allPossibleTimeSlots[startTimeIndex]+ ":00-08:00"
    currentSlotEnd := extractedEndsAtDate + "T" + allPossibleTimeSlots[endTimeIndex]+ ":00-08:00"

		if _, ok := bookedApptWithinGivenDateRange[currentSlotStart]; !ok {
			newSlot["starts_at"] = currentSlotStart
			newSlot["ends_at"] = currentSlotEnd
			openAppointmentSlots = append(openAppointmentSlots, newSlot)
		}
		startTimeIndex++
		endTimeIndex++
	}
	return openAppointmentSlots
}

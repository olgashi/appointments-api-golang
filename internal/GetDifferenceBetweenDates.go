package internal

import (
	"time"
	"fmt"
)

func GetDifferenceBetweenDatesInMinutes(startDateTime, endDateTime string) float64 {
	startDateTimeParsed, err1 := time.Parse(time.RFC3339, startDateTime)
	endDateTimeParsed, err2 := time.Parse(time.RFC3339, endDateTime)

	if err1 != nil || err2 != nil {
		fmt.Println("Could not parse time:", err1, err2)
	}

	return endDateTimeParsed.Sub(startDateTimeParsed).Minutes()
}
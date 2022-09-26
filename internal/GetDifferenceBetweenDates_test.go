package internal

import (
	"testing"
)

func TestGetDifferenceBetweenDatesInMinutes_DifferenceIsThirtyMinutes(t *testing.T) {
	differenceInMinutes := GetDifferenceBetweenDatesInMinutes("2019-01-24T09:00:00-08:00", "2019-01-24T09:30:00-08:00")
	if differenceInMinutes != 30 {
		t.Fatalf(`Should return difference between two dates in minutes, but returned %v`, differenceInMinutes)
	}
}

func TestGetDifferenceBetweenDatesInMinutes_DifferenceIsGreaterThanThirtyMinutes(t *testing.T) {
	differenceInMinutes := GetDifferenceBetweenDatesInMinutes("2019-01-24T09:00:00-08:00", "2019-01-24T10:30:00-08:00")
	if differenceInMinutes != 90 {
		t.Fatalf(`Should return difference between two dates in minutes, but returned %v`, differenceInMinutes)
	}
}
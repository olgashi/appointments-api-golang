package internal

import (
	"testing"
)

func TestIsWithinBusinessHours(t *testing.T) {
	dateIsWithinBusinessHours := IsWithinBusinessHours("2019-01-24T10:30:00-08:00", "2019-01-24T10:30:00-08:00")
	if dateIsWithinBusinessHours != true {
		t.Fatalf(`Should return true but returned %v`, dateIsWithinBusinessHours)
	}
}

func TestIsWithinBusinessHours_OutsideOfBusinessHours(t *testing.T) {
	dateIsWithinBusinessHours := IsWithinBusinessHours("2019-01-24T19:30:00-08:00", "2019-01-24T20:00:00-08:00")
	if dateIsWithinBusinessHours != false {
		t.Fatalf(`Should return false but returned %v`, dateIsWithinBusinessHours)
	}
}
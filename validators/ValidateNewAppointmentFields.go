package validators

import (
	"appointments-api/models"
)

func ValidateNewAppointmentFields(appt models.SubmittedAppointment) []string {
	var errors[] string
	if appt.Ends_at == "" {
		errors = append(errors, "'ends_at' is missing")
	} else if appt.Starts_at == "" {
		errors = append(errors, "'starts_at' is missing")
	} else if appt.User_id == 0 {
		errors = append(errors, "'user_id' is missing")
	} else if appt.Trainer_id == 0 {
		errors = append(errors, "'trainer_id' is missing")
	}

	return errors
}
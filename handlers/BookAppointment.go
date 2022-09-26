package handlers

import (
	"net/http"
	"strconv"

	"appointments-api/internal"
	"appointments-api/mocks"
	"appointments-api/models"
	"appointments-api/validators"

	"github.com/gin-gonic/gin"
)

// To book appointments
func BookAppointment(c *gin.Context) {
	const appointmentDuration = 30
	payload := models.SubmittedAppointment{}
	c.BindJSON(&payload)

	errors := validators.ValidateNewAppointmentFields(payload)
	if len(errors) > 0 {
		c.IndentedJSON(http.StatusBadRequest, errors)
		return
	}
	
	if (!internal.TrainerExists(payload.Trainer_id, mocks.Trainers)) {
		c.IndentedJSON(http.StatusNotFound, "The trainer for the provided trainer id does not exist")
	  return
	}
	
	if (!internal.UserExists(payload.User_id, mocks.Users)) {
		c.IndentedJSON(http.StatusNotFound, "The user for the provided user id does not exist")
	  return
	}

	if (!internal.IsWithinBusinessHours(payload.Starts_at, payload.Ends_at)) {
		c.IndentedJSON(http.StatusBadRequest, "The appointment is outside business hours")
		return
	}
	
	if (internal.IsAlreadyBooked(payload.Starts_at, payload.Ends_at, strconv.Itoa(payload.Trainer_id))) {
		c.IndentedJSON(http.StatusBadRequest, "Appointment for that time slot and for the provided trainer id already booked")
	  return
	}

	if (internal.GetDifferenceBetweenDatesInMinutes(payload.Starts_at, payload.Ends_at) != appointmentDuration) {
		c.IndentedJSON(http.StatusBadRequest, "Appointment length must be exactly 30 minutes")
	  return
	}

	newApptId := internal.GenerateAppointmentId(mocks.Appointments)
	newAppointment := internal.CreateNewAppointment(payload.Ends_at, newApptId, payload.User_id, payload.Starts_at, payload.Trainer_id)
	mocks.Appointments = append(mocks.Appointments, *newAppointment)
	
	c.JSON(http.StatusCreated, gin.H{"data": 
	gin.H{
		"message": "appointment created successfully",
		"appointment_id": newApptId,
	  },
  })
}
package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"appointments-api/internal"
	"appointments-api/mocks"
	"strconv"
)

// Get scheduled/booked appointments for a specific trainer id
func GetScheduledAppointments(c *gin.Context) {
	id := c.Param("trainer_id")

	idStr, _ := strconv.Atoi(id)
	if (!internal.TrainerExists(idStr, mocks.Trainers)) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "trainer does not exist"})
		return
	}

	bookedAppointments := internal.FilterAppointmentsByTrainerId(id, mocks.Appointments)

	c.JSON(http.StatusOK, gin.H{"data": 
	gin.H{
		"appointments": bookedAppointments,
		"trainer_id": id,
	},
	})
}
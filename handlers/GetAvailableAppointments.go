package handlers

import (
	"appointments-api/internal"
	"appointments-api/mocks"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Get available/open appointment slots for a specific trainer id
func GetAvailableAppointments(c *gin.Context) {
	id := c.Param("trainer_id")
	startDate := c.Query("starts_at")
	endDate := c.Query("ends_at")
	startDate = startDate + "T08:00:00-08:00"
	endDate = endDate + "T17:00:00-08:00"

	idStr, _ := strconv.Atoi(id)
	
	if (!internal.TrainerExists(idStr, mocks.Trainers)) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "trainer does not exist"})
		return
	}

	openAppointments := internal.GetAllAvailableAppointmentsForDateRange(startDate, endDate, id)


	c.IndentedJSON(http.StatusOK, openAppointments)
}
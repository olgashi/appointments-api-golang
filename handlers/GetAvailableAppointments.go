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

	if !internal.TrainerExists(idStr, mocks.Trainers) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "trainer does not exist"})
		return
	}

	openAppointments := internal.GetAllAvailableAppointmentsForDateRange(startDate, endDate, id)
	strId, _ := strconv.Atoi(id)

	type customResponse struct {
		appointments []map[string]string
		trainer_id   int
	}

	var openAppointmentsResponce customResponse

	openAppointmentsResponce.appointments = openAppointments
	openAppointmentsResponce.trainer_id = strId

	c.JSON(http.StatusOK, gin.H{"data": 
	gin.H{
		"appointments": openAppointments,
		"trainer_id": id,
	},
	})
}

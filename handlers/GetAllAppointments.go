package handlers
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"appointments-api/mocks"
)

// To get all existing appointments
func GetAllAppointments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": 
	gin.H{
		"appointments": mocks.Appointments,
		},
	})
}
package main

import (
	"appointments-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/v1/appointments/all", handlers.GetAllAppointments)
	router.POST("/v1/appointments", handlers.BookAppointment)
	router.GET("/v1/appointments/scheduled/:trainer_id", handlers.GetScheduledAppointments)
	router.GET("/v1/appointments/available/:trainer_id", handlers.GetAvailableAppointments)
	router.Run("localhost:8080")
}
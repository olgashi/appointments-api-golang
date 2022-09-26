package main

import (
	"appointments-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/appointments/all", handlers.GetAllAppointments)
	router.POST("/appointments", handlers.BookAppointment)
	router.GET("/appointments/scheduled/:trainer_id", handlers.GetScheduledAppointments)
	router.GET("/appointments/available/:trainer_id", handlers.GetAvailableAppointments)
	router.Run("localhost:8080")
}
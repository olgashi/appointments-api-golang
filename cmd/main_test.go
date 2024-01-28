package main

import (
	"appointments-api/handlers"
	"appointments-api/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetScheduledAppointments_TrainerExists(t *testing.T) {
	r := SetUpRouter()
	r.GET("/v1/appointments/scheduled/:trainer_id", handlers.GetScheduledAppointments)

	req, _ := http.NewRequest("GET", "/v1/appointments/scheduled/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetScheduledAppointments_TrainerDoesntExist(t *testing.T) {
	r := SetUpRouter()
	r.GET("/v1/appointments/scheduled/:trainer_id", handlers.GetScheduledAppointments)

	req, _ := http.NewRequest("GET", "/v1/appointments/scheduled/100", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBookAppointment(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2020-01-26T12:30:00-08:00",
		User_id:    8,
		Starts_at:  "2020-01-26T12:00:00-08:00",
		Trainer_id: 3,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestBookAppointment_OutsideBusinessHours(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2020-01-26T19:30:00-08:00",
		User_id:    8,
		Starts_at:  "2020-01-26T19:00:00-08:00",
		Trainer_id: 3,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBookAppointment_Doublebooked(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2019-01-26T10:30:00-08:00",
		User_id:    3,
		Starts_at:  "2019-01-26T10:00:00-08:00",
		Trainer_id: 1,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBookAppointment_InvalidDuration(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2020-01-26T12:30:00-08:00",
		User_id:    3,
		Starts_at:  "2020-01-26T10:00:00-08:00",
		Trainer_id: 1,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBookAppointment_TrainerDoesntExist(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2020-01-26T12:30:00-08:00",
		User_id:    8,
		Starts_at:  "2020-01-26T12:00:00-08:00",
		Trainer_id: 31,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBookAppointment_UserDoesntExist(t *testing.T) {
	r := SetUpRouter()
	r.POST("/v1/appointments", handlers.BookAppointment)

	appointment := models.SubmittedAppointment{
		Ends_at:    "2020-01-26T12:30:00-08:00",
		User_id:    88,
		Starts_at:  "2020-01-26T12:00:00-08:00",
		Trainer_id: 3,
	}
	jsonValue, _ := json.Marshal(appointment)
	req, _ := http.NewRequest("POST", "/v1/appointments", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGetAvailableAppointments(t *testing.T) {
	r := SetUpRouter()
	r.GET("/v1/appointments/available/:trainer_id", handlers.GetAvailableAppointments)

	req, _ := http.NewRequest("GET", "/v1/appointments/available/1?starts_at=2018-01-24&ends_at=2018-01-24", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAvailableAppointments_TrainerDoesNotexist(t *testing.T) {
	r := SetUpRouter()
	r.GET("/v1/appointments/available/:trainer_id", handlers.GetAvailableAppointments)

	req, _ := http.NewRequest("GET", "/v1/appointments/available/103?starts_at=2018-01-24&ends_at=2018-01-24", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

package internal

import (
	"appointments-api/models"
)

func UserExists(id int, users []models.User) bool {
  for _, user := range users {
		if user.User_id == id {
			return true
		}
	}
	return false
}
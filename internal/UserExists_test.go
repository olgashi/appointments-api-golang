package internal

import (
	"testing"
	"appointments-api/models"
)
var users = []models.User{
	{
		User_id: 1,
	},
	{
		User_id: 2,
	},
	{
		User_id: 3,
	},
	{
		User_id: 4,
	},
	{
		User_id: 5,
	},
}

func TestUserExists(t *testing.T) {
	userId := 3
	userExistsInList := UserExists(userId, users)

	if userExistsInList != true {
		t.Fatalf(`Should return true for user id %v, but returned %v`, userId, userExistsInList)
	}
}

func TestUserExists_UserDoesntExist(t *testing.T) {
	userId := 1001
	userExistsInList := UserExists(userId, users)

	if userExistsInList != false {
		t.Fatalf(`Should return false for user id %v, but returned %v`, userId, userExistsInList)
	}
}
package mocks

import 	models "appointments-api/models"

// Provided appointment data
var Appointments = []models.Appointment{
  {
		Ended_at: "2019-01-24T09:30:00-08:00",
        ID: 1,
        User_id: 1,
        Started_at: "2019-01-24T09:00:00-08:00",
        Trainer_id: 1,
    },
    {
        Ended_at: "2019-01-24T10:30:00-08:00",
        ID: 2,
        User_id: 2,
        Started_at: "2019-01-24T10:00:00-08:00",
        Trainer_id: 1,
    },
    {
        Ended_at: "2019-01-25T10:30:00-08:00",
        ID: 3,
        User_id: 3,
        Started_at: "2019-01-25T10:00:00-08:00",
        Trainer_id: 1,
    },
    {
        Ended_at: "2019-01-25T11:00:00-08:00",
        ID: 4,
        User_id: 4,
        Started_at: "2019-01-25T10:30:00-08:00",
        Trainer_id: 1,
    },
    {
        Ended_at: "2019-01-26T10:30:00-08:00",
        ID: 5,
        User_id: 5,
        Started_at: "2019-01-26T10:00:00-08:00",
        Trainer_id: 1,
    },
    {
        Ended_at: "2019-01-24T09:30:00-08:00",
        ID: 6,
        User_id: 6,
        Started_at: "2019-01-24T09:00:00-08:00",
        Trainer_id: 2,
    },
    {
        Ended_at: "2019-01-26T10:30:00-08:00",
        ID: 7,
        User_id: 7,
        Started_at: "2019-01-26T10:00:00-08:00",
        Trainer_id: 2,
    },
    {
        Ended_at: "2019-01-26T12:30:00-08:00",
        ID: 8,
        User_id: 8,
        Started_at: "2019-01-26T12:00:00-08:00",
        Trainer_id: 3,
    },
    {
        Ended_at: "2019-01-26T14:00:00-08:00",
        ID: 9,
        User_id: 9,
        Started_at: "2019-01-26T13:00:00-08:00",
        Trainer_id: 3,
    },
    {
        Ended_at: "2019-01-26T14:30:00-08:00",
        ID: 10,
        User_id: 10,
        Started_at: "2019-01-26T14:00:00-08:00",
        Trainer_id: 3,
    },
}

// Added the following data, based on provided data
var Users = []models.User{
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
	{
		User_id: 6,
	},
	{
		User_id: 7,
	},
	{
		User_id: 8,
	},
	{
		User_id: 9,
	},
	{
		User_id: 10,
	},
}

var Trainers = []models.Trainer{
	{
		Trainer_id: 1,
	},
	{
		Trainer_id: 2,
	},
	{
		Trainer_id: 3,
	},
}
## Appointment Scheduling API

## To run the Application:

`git clone https://github.com/olgashi/appointments-api-golang.git`

`cd appointments-api-golang/cmd`

`go run .`

The application is now running on `localhost:8080`

Feel free to try any of the endpoints by using `localhost:8080` as the base url.

Example: `http://localhost:8080/v1/appointments/all`


## Tests
- To run Unit Tests for REST API endpoints
From projects root directory 

`cd internal`

`go test`


- To run unit tests for logic and utility functions
From projects root directory 

`cd cmd`

`go test`


## Folder Structure
```
├── cmd
│   ├── main_test.go
│   ├── main.go - `entry point`
├── handlers - `handler functions for each of the defined routes`
├── internal - `logic and utility functions and unit tests`
├── mocks - `application data`
├── models - `data schemas`
└── validators - `validator functions`
```

### `GET /v1/appointments/available/:trainer_id`

 __Get a list of available appointment start and end times for a trainer between two dates__

 ### Query and Path Parameters:

  - `trainer_id` (requred)

  Id of the trainer you wish to view all the available appointments for

  - `starts_at` (required) 
  
  Appointment start datetime, should be specified as YYYY-MM-DD, example: 2019-01-24

  - `ends_at` (required)

  Appointment end datetime, should be specified as YYYY-MM-DD, example: 2019-01-25


  Example: 
  `/v1/appointments/available/1?starts_at=2019-01-24&ends_at=2019-01-24`

 Returns list of available appointment times between 8 am on 5 pm on 2019-01-24

Success Status code: `200 OK`

Example Successful Response:
 ```json
{
  "data": {
    "appointments": [
      {
        "ends_at": "2019-01-24T08:30:00-08:00",
        "starts_at": "2019-01-24T08:00:00-08:00"
      },
      {
        "ends_at": "2019-01-24T12:00:00-08:00",
        "starts_at": "2019-01-24T11:30:00-08:00"
      },
      {
        "ends_at": "2019-01-24T13:30:00-08:00",
        "starts_at": "2019-01-24T13:00:00-08:00"
      },
      {
        "ends_at": "2019-01-24T15:00:00-08:00",
        "starts_at": "2019-01-24T14:30:00-08:00"
      },
      {
        "ends_at": "2019-01-24T15:30:00-08:00",
        "starts_at": "2019-01-24T15:00:00-08:00"
      },
      {
        "ends_at": "2019-01-24T16:00:00-08:00",
        "starts_at": "2019-01-24T15:30:00-08:00"
      },
      {
        "ends_at": "2019-01-24T16:30:00-08:00",
        "starts_at": "2019-01-24T16:00:00-08:00"
      },
      {
        "ends_at": "2019-01-24T17:00:00-08:00",
        "starts_at": "2019-01-24T16:30:00-08:00"
      }
    ],
    "trainer_id": "1"
  }
}
```



### `POST /v1/appointments`

__Create a new appointment (as JSON)__

### Request Body Parameters:

 - `trainer_id` (required)

Id of the trainer you wish to book the appointment for

 - `user_id` (required)

Id of the user booking the appointment

 - `starts_at` (required)

 The start datetime of the appointment in the following format: 2019-01-24T10:00:00-08:00

 - `ends_at` (required)
  The end datetime of the appointment in the following format: 2019-01-24T10:00:00-08:00

Success Status code: `201 CREATED`

Example Successful Response:

```json
{
  "data":
  {
    "appointment_id": 11,
    "message": "appointment created successfully"
  }
}
```



### `GET /v1/appointments/scheduled/:trainer_id`
Get a list of scheduled appointments for a trainer

### Parameters:

 - `trainer_id` (required)
Id of the trainer you wish to view all the scheduled appointments for

Example:
`v1/appointments/scheduled/1`

Success Status code: `200 OK`

Example Successful Response:

```json
{
  "data": {
    "appointments": [
      {
        "ended_at": "2019-01-24T09:30:00-08:00",
        "id": 1,
        "user_id": 1,
        "started_at": "2019-01-24T09:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-24T10:30:00-08:00",
        "id": 2,
        "user_id": 2,
        "started_at": "2019-01-24T10:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-25T10:30:00-08:00",
        "id": 3,
        "user_id": 3,
        "started_at": "2019-01-25T10:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-25T11:00:00-08:00",
        "id": 4,
        "user_id": 4,
        "started_at": "2019-01-25T10:30:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-26T10:30:00-08:00",
        "id": 5,
        "user_id": 5,
        "started_at": "2019-01-26T10:00:00-08:00",
        "trainer_id": 1
      }
    ],
    "trainer_id": "1"
  }
}
```



### `GET /v1/appointments/all`
Get a list of all scheduled appointments

Example Response Body:

```json
{
  "data": {
    "appointments": [
      {
        "ended_at": "2019-01-24T09:30:00-08:00",
        "id": 1,
        "user_id": 1,
        "started_at": "2019-01-24T09:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-24T10:30:00-08:00",
        "id": 2,
        "user_id": 2,
        "started_at": "2019-01-24T10:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-25T10:30:00-08:00",
        "id": 3,
        "user_id": 3,
        "started_at": "2019-01-25T10:00:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-25T11:00:00-08:00",
        "id": 4,
        "user_id": 4,
        "started_at": "2019-01-25T10:30:00-08:00",
        "trainer_id": 1
      },
      {
        "ended_at": "2019-01-26T10:30:00-08:00",
        "id": 5,
        "user_id": 5,
        "started_at": "2019-01-26T10:00:00-08:00",
        "trainer_id": 1
      },
    ]
  }
  ```

# Time Tracker API

This project is a REST API for a time tracking system. 
The application supports user and task creation, time tracking, task start and end operations, and checking the current state of tasks. 
All operations are logged to the console.

## Features

- **Creating a new user**
- **Creating a new task**
- **Ending a task**
- **Checking the current state of tasks**

## Architecture

The project follows the principles of pure architecture, separating business logic, infrastructure and interfaces. 
A service provider is used for dependency management and dependency injection.

## Technologies Used

- **Golang**: Backend service implemented in Go.
- **Docker**: Containerization for easy deployment and scalability.
- **Gin**: HTTP web framework for Go.
- **PostgreSQL**: Relational database management system.
- **Swaggo**: Swagger for Go, used to auto-generate RESTful API documentation.

## Setup Instructions

Follow these steps to set up and run the Time Tracker API locally:

### Clone the repository

```bash
git clone https://github.com/PabloPerdolie/time-tracker
cd time-tracker
```
### Launching the application

```bash
make run
```

The service will start running on http://localhost:8080/.

#### To clean up:
```bash
make clean
```

## Endpoints

You can use Postman to test this API.

#### Create New User

- **URL**: `/users`
- **Method**: `POST`
- **Request Body**:
    ````json
    {
      "passportNumber": "1234 567890"
    }
    ````
- **Response**:
  - `201 Created` New account created.
    - Sample response:
       ````json
       {
          "address": "string",
          "created_at": "string",
          "id": 0,
          "name": "string",
          "passport_number": "string",
          "passport_series": "string",
          "patronymic": "string",
          "surname": "string",
          "updated_at": "string"
       }
       ````
  - `400 Bad Request`: Invalid request.
  - `500 Internal Server Error`: Server error.

#### Create New Task

- **URL**: `/tasks/start/{id}`
- **Method**: `POST`
- **Request Body**:
    ````json
    {
      "description": "Task n"
    }
    ````
- **Response**:
  - `200 OK` New task created.
    - Sample response:
       ````json
       {
          "created_at": "string",
          "description": "string",
          "duration": "string",
          "end_time": "string",
          "id": 0,
          "start_time": "string",
          "updated_at": "string",
          "user_id": 0
       }
       ````
  - `400 Bad Request`: Invalid request.
  - `500 Internal Server Error`: Server error.

#### End Task

- **URL**: `/tasks/end/{taskId}`
- **Method**: `POST`
- **Response**:
  - `200 OK` Task ended.
    - Sample response:
         ````json
         {
            "created_at": "string",
            "description": "string",
            "duration": "string",
            "end_time": "string",
            "id": 0,
            "start_time": "string",
            "updated_at": "string",
            "user_id": 0
         }
         ````
  - `400 Bad Request`: Invalid request.
  - `500 Internal Server Error`: Server error.

#### Get Tasks by User with period

- **URL**: `/tasks/user/{id}`
- **Method**: `GET`
- **Response**:
  - `200 OK`.
    - Sample response:
       ````json
       [
        {
            "created_at": "string",
            "description": "string",
            "duration": "string",
            "end_time": "string",
            "id": 0,
            "start_time": "string",
            "updated_at": "string",
            "user_id": 0
         }
       ]
       ````

<center>Thanks for checking out my service.</center>

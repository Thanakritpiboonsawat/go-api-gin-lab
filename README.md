# Student API with Gin and SQLite

## How to Run

1. Install Go ก่อน
2. Clone ลง repository
3. Run:ด้วยคําสั่ง go run main.go

Base URL: http://localhost:8080

## Project Structure

- main.go
- config/
- models/
- repositories/
- services/
- handlers/

## API Endpoints

### Get all students
GET /students

### Get student by ID
GET /students/:id

### Create student
POST /students 

Example JSON:
{
  "name": "John Doe",
  "age": 20,
  "email": "john@example.com"
}

### Update student
PUT /students/:id

### Delete student
DELETE /students/:id


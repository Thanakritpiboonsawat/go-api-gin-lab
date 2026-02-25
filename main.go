package main

import (
	"github.com/gin-gonic/gin"

	"example.com/student-gin-api/config"
	"example.com/student-gin-api/handlers"
	"example.com/student-gin-api/repositories"
	"example.com/student-gin-api/services"
)

func main() {
	db := config.InitDB()

	repo := &repositories.StudentRepository{DB: db}
	service := &services.StudentService{Repo: repo}
	handler := &handlers.StudentHandler{Service: service}

	r := gin.Default()

	r.GET("/students", handler.GetStudents)
	r.GET("/students/:id", handler.GetStudentByID)
	r.POST("/students", handler.CreateStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)

	r.Run(":8080")
}

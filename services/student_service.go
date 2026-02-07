package services

import (
	"example.com/student-gin-api/models"
	"example.com/student-gin-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}
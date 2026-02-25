package services

import (
	"database/sql"
	"errors"

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

func (s *StudentService) UpdateStudent(id string, student models.Student) (*models.Student, error) {
	return s.Repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}

// Optional: Custom error helper
func IsNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}

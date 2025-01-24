package repository

import (
	"student-management/pkg/model"
)

// StudentRepository is an interface that defines student repository methods
type StudentRepository interface {
	GetAllStudents() ([]model.Student, error)
	AddStudent(student model.Student) (int, error)
	UpdateStudent(studentId int, student model.Student) error
	DeleteStudent(studentId int) error
	GetStudentByID(studentId int) (model.Student, error)
}

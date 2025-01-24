package service

import (
	"student-management/pkg/model"
	"student-management/pkg/repository"
	"time"
)

// StudentService is an interface that defines student-related service methods
type StudentService interface {
	GetStudents() ([]model.Student, error)
	AddStudent(student model.Student) (int, error)
	UpdateStudent(studentID int, student model.Student) error
	DeleteStudent(studentID int) error
	GetStudentByID(studentID int) (model.Student, error)
}

// StudentServiceImpl implements StudentService and contains business logic
type StudentServiceImpl struct {
	StudentRepo repository.StudentRepository
}

// NewStudentService creates a new Student service
func NewStudentService(sRepo repository.StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{StudentRepo: sRepo}
}

// GetStudents fetches Students from the repository
func (stdService *StudentServiceImpl) GetStudents() ([]model.Student, error) {
	Students, err := stdService.StudentRepo.GetAllStudents()
	if err != nil {
		return nil, err
	}
	return Students, nil
}

func (stdService *StudentServiceImpl) AddStudent(student model.Student) (int, error) {
	student.EnrollmentDate = time.Now().Format("2006-01-02")
	student.Status = "Active"
	id, err := stdService.StudentRepo.AddStudent(student)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (stdService *StudentServiceImpl) UpdateStudent(studentID int, student model.Student) error {
	return stdService.StudentRepo.UpdateStudent(studentID, student)
}

func (stdService *StudentServiceImpl) DeleteStudent(studentID int) error {
	return stdService.StudentRepo.DeleteStudent(studentID)
}

func (stdService *StudentServiceImpl) GetStudentByID(studentID int) (model.Student, error) {
	return stdService.StudentRepo.GetStudentByID(studentID)
}

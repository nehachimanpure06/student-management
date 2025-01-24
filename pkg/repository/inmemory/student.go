package inmemory

import (
	"errors"
	"student-management/pkg/model"
)

var InMemoryStudentData = []model.Student{
	{ID: 1, FirstName: "neha", LastName: "chimanpure", Email: "neha@example.com"},
	{ID: 2, FirstName: "aditi", LastName: "kasar", Email: "aditi@example.com"},
}

// InMemoryStudentRepository implements StudentRepository for in-memory storage
type InMemoryStudentRepository struct {
	Students []model.Student
}

// NewStudentRepository creates a new in-memory Student repository
func NewStudentRepository() *InMemoryStudentRepository {
	return &InMemoryStudentRepository{
		Students: InMemoryStudentData,
	}
}

// GetAllStudents returns all Students in the in-memory store
func (repo *InMemoryStudentRepository) GetAllStudents() ([]model.Student, error) {
	return repo.Students, nil
}

func (repo *InMemoryStudentRepository) AddStudent(student model.Student) (int, error) {
	// Simulating ID assignment
	student.ID = len(repo.Students) + 1

	repo.Students = append(repo.Students, student)
	return student.ID, nil
}

func (repo *InMemoryStudentRepository) GetStudentByID(studentId int) (model.Student, error) {
	for _, student := range repo.Students {
		if student.ID == studentId {
			return student, nil
		}
	}
	return model.Student{}, errors.New("student with given id does not exists")
}

func (repo *InMemoryStudentRepository) UpdateStudent(studentId int, updateStudent model.Student) error {
	// When you use for _, student := range repo.Students, you're working with a copy of the student
	// from the slice. Changes to that copy don’t modify the original slice.
	for i := range repo.Students {
		if repo.Students[i].ID == studentId {
			repo.Students[i].FirstName = updateStudent.FirstName
			repo.Students[i].LastName = updateStudent.LastName
			repo.Students[i].Email = updateStudent.Email
			repo.Students[i].Phone = updateStudent.Phone
			repo.Students[i].DateOfBirth = updateStudent.DateOfBirth
			repo.Students[i].EnrollmentDate = updateStudent.EnrollmentDate
			repo.Students[i].Status = updateStudent.Status
			return nil
		}
	}
	return errors.New("student with given id does not exist")
}

func (repo *InMemoryStudentRepository) DeleteStudent(studentId int) error {
	for i, student := range repo.Students {
		if student.ID == studentId {
			// Remove the student from the slice by concatenating the two parts
			repo.Students = append(repo.Students[:i], repo.Students[i+1:]...)
			return nil
		}
	}
	return errors.New("student with given id does not exists")
}

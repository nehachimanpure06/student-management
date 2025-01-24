package service

import (
	"student-management/pkg/model"
	"student-management/pkg/repository"
)

// CourseService is an interface that defines course-related service methods
type CourseService interface {
	GetCourses() ([]model.Course, error)
	AddCourse(course model.Course) (int, error)
	UpdateCourse(courseID int, course model.Course) error
	DeleteCourse(courseID int) error
	GetCourseByID(courseID int) (model.Course, error)
}

// CourseServiceImpl implements CourseService and contains business logic
type CourseServiceImpl struct {
	CourseRepo repository.CourseRepository
}

// NewCourseService creates a new Course service
func NewCourseService(cRepo repository.CourseRepository) *CourseServiceImpl {
	return &CourseServiceImpl{CourseRepo: cRepo}
}

// GetCourses fetches Courses from the repository
func (cService *CourseServiceImpl) GetCourses() ([]model.Course, error) {
	Courses, err := cService.CourseRepo.GetAllCourses()
	if err != nil {
		return nil, err
	}
	return Courses, nil
}

func (cService *CourseServiceImpl) AddCourse(course model.Course) (int, error) {
	id, err := cService.CourseRepo.AddCourse(course)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (cService *CourseServiceImpl) UpdateCourse(courseID int, course model.Course) error {
	return cService.CourseRepo.UpdateCourse(courseID, course)
}

func (cService *CourseServiceImpl) DeleteCourse(courseID int) error {
	return cService.CourseRepo.DeleteCourse(courseID)
}

func (cService *CourseServiceImpl) GetCourseByID(courseID int) (model.Course, error) {
	return cService.CourseRepo.GetCourseByID(courseID)
}

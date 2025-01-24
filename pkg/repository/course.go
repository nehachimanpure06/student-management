package repository

import (
	"student-management/pkg/model"
)

// CourseRepository is an interface that defines course repository methods
type CourseRepository interface {
	GetAllCourses() ([]model.Course, error)
	AddCourse(course model.Course) (int, error)
	UpdateCourse(courseId int, course model.Course) error
	DeleteCourse(courseId int) error
	GetCourseByID(courseId int) (model.Course, error)
}

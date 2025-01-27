package inmemory

import (
	"errors"
	"student-management/pkg/model"
)

var InMemoryCourseData = []model.Course{
	{
		ID:             1,
		Name:           "golang",
		Description:    "this course covers basic go concepts to start go developemnt joury",
		Credits:        100,
		Instructor:     "neha chimanpure",
		Schedule:       "monday 01:00-02:00",
		Capacity:       50,
		AvailableSeats: 50,
	},
}

// InMemoryCourseRepository implements CourseRepository for in-memory storage
type InMemoryCourseRepository struct {
	Courses []model.Course
}

// NewCourseRepository creates a new in-memory Course repository
func NewCourseRepository() *InMemoryCourseRepository {
	return &InMemoryCourseRepository{
		Courses: InMemoryCourseData,
	}
}

// GetAllCourses returns all Courses in the in-memory store
func (repo *InMemoryCourseRepository) GetAllCourses() ([]model.Course, error) {
	return repo.Courses, nil
}

func (repo *InMemoryCourseRepository) AddCourse(course model.Course) (int, error) {
	// Simulating ID assignment
	course.ID = len(repo.Courses) + 1

	repo.Courses = append(repo.Courses, course)
	return course.ID, nil
}

func (repo *InMemoryCourseRepository) GetCourseByID(courseId int) (model.Course, error) {
	for _, course := range repo.Courses {
		if course.ID == courseId {
			return course, nil
		}
	}
	return model.Course{}, errors.New("course with given id does not exists")
}

func (repo *InMemoryCourseRepository) UpdateCourse(courseId int, updateCourse model.Course) error {
	// When you use for _, course := range repo.Courses, you're working with a copy of the course
	// from the slice. Changes to that copy don’t modify the original slice.
	for i := range repo.Courses {
		if repo.Courses[i].ID == courseId {
			repo.Courses[i].Name = updateCourse.Name
			repo.Courses[i].Description = updateCourse.Description
			repo.Courses[i].Credits = updateCourse.Credits
			repo.Courses[i].Instructor = updateCourse.Instructor
			repo.Courses[i].Schedule = updateCourse.Schedule
			repo.Courses[i].Capacity = updateCourse.Capacity
			repo.Courses[i].AvailableSeats = updateCourse.AvailableSeats
			return nil
		}
	}
	return errors.New("course with given id does not exist")
}

func (repo *InMemoryCourseRepository) DeleteCourse(courseId int) error {
	for i, course := range repo.Courses {
		if course.ID == courseId {
			// Remove the course from the slice by concatenating the two parts
			repo.Courses = append(repo.Courses[:i], repo.Courses[i+1:]...)
			return nil
		}
	}
	return errors.New("course with given id does not exists")
}

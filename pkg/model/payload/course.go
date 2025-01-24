package payload

import "student-management/pkg/model"

type CourseRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Credits     int    `json:"credits" validate:"required"`
	Instructor  string `json:"instructor" validate:"required"`
	Schedule    string `json:"schedule" validate:"required"`
}

func ToCourseModel(courseRequest CourseRequest) model.Course {
	return model.Course{
		Name:        courseRequest.Name,
		Description: courseRequest.Description,
		Credits:     courseRequest.Credits,
		Instructor:  courseRequest.Instructor,
		Schedule:    courseRequest.Schedule,
	}
}

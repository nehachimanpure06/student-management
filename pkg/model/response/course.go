package response

import "student-management/pkg/model"

type CourseResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Credits        int    `json:"credits"`
	Instructor     string `json:"instructor"`
	Schedule       string `json:"schedule"`
	Capacity       int    `json:"capacity"`
	AvailableSeats int    `json:"available_seats"`
}

func ToCourseResponse(course model.Course) CourseResponse {
	return CourseResponse{
		ID:             course.ID,
		Name:           course.Name,
		Description:    course.Description,
		Credits:        course.Credits,
		Instructor:     course.Instructor,
		Schedule:       course.Schedule,
		Capacity:       course.Capacity,
		AvailableSeats: course.AvailableSeats,
	}
}

func ToCourseListResponse(courses []model.Course) []CourseResponse {
	var responseList []CourseResponse
	for _, course := range courses {
		response := ToCourseResponse(course)
		responseList = append(responseList, response)
	}
	return responseList
}

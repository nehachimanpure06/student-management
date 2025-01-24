package response

import "student-management/pkg/model"

type StudentResponse struct {
	ID             int    `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	DateOfBirth    string `json:"date_of_birth"`
	EnrollmentDate string `json:"enrollment_date"`
	Status         string `json:"status"`
}

func ToStudentResponse(student model.Student) StudentResponse {
	return StudentResponse{
		ID:             student.ID,
		FirstName:      student.FirstName,
		LastName:       student.LastName,
		Email:          student.Email,
		Phone:          student.Phone,
		DateOfBirth:    student.DateOfBirth,
		EnrollmentDate: student.EnrollmentDate,
		Status:         student.Status,
	}
}

func ToStudentListResponse(students []model.Student) []StudentResponse {
	var responseList []StudentResponse
	for _, student := range students {
		response := ToStudentResponse(student)
		responseList = append(responseList, response)
	}
	return responseList
}

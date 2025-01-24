package payload

import "student-management/pkg/model"

type StudentRequest struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Phone       string `json:"phone" validate:"required,len=10,numeric"`
	DateOfBirth string `json:"date_of_birth" validate:"required"`
}

func ToStudentModel(studentReq StudentRequest) model.Student {
	return model.Student{
		FirstName:   studentReq.FirstName,
		LastName:    studentReq.LastName,
		Email:       studentReq.Email,
		Phone:       studentReq.Phone,
		DateOfBirth: studentReq.DateOfBirth,
	}
}

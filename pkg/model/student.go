package model

type Student struct {
	ID             int
	FirstName      string
	LastName       string
	Email          string
	Phone          string
	DateOfBirth    string
	EnrollmentDate string
	Status         string // Active, Graduated, Dropped
	CourseIDs      []int
}

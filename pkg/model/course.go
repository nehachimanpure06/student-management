package model

type Course struct {
	ID             int
	Name           string
	Description    string
	Credits        int
	Instructor     string
	Schedule       string
	Capacity       int
	AvailableSeats int
}

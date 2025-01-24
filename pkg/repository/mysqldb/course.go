package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"student-management/pkg/model"
)

// MysqlCourseRepository implements CourseRepository using mysql database
type MysqlCourseRepository struct {
	DB *sql.DB
}

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

// NewMysqlCourseRepository creates a new PostgresCourseRepository instance
func NewMysqlCourseRepository(db *sql.DB) (*MysqlCourseRepository, error) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS courses(
	id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL,
    credits INT,
    instructor VARCHAR(100),
    schedule varchar(100),
    capacity int DEFAULT 50,
    available_seats int DEFAULT 50)`)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating table courses : %v", err.Error())
	}
	return &MysqlCourseRepository{DB: db}, nil
}

// GetAllCourses retrieves all courses from the mysql database
func (repo *MysqlCourseRepository) GetAllCourses() ([]model.Course, error) {
	var courses []model.Course
	rows, err := repo.DB.Query("SELECT id, name, description, credits, instructor, schedule, capacity, available_seats FROM courses")
	if err != nil {
		return nil, errors.New("error occured while getting course data : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var course model.Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.Credits, &course.Instructor, &course.Schedule, &course.Capacity, &course.AvailableSeats)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

// AddCourse adds a new course to the mysql database
func (repo *MysqlCourseRepository) AddCourse(course model.Course) (int, error) {
	query := `INSERT INTO courses(name, description, credits, instructor, schedule)
	VALUES (?, ?, ?, ?, ?)`

	result, err := repo.DB.Exec(query, course.Name, course.Description, course.Credits, course.Instructor, course.Schedule)
	if err != nil {
		return 0, err
	}

	// Get the last inserted ID
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

// UpdateCourse updates an existing course in the mysql database
func (repo *MysqlCourseRepository) UpdateCourse(courseID int, course model.Course) error {
	query := `UPDATE courses SET name = ?, description = ?, credits = ?, instructor = ?, schedule = ? WHERE id = ?`

	_, err := repo.DB.Exec(query, course.Name, course.Description, course.Credits, course.Instructor, course.Schedule, courseID)
	if err != nil {
		return fmt.Errorf("could not update course: %v", err)
	}

	return nil
}

// DeleteCourse deletes a course from the mysql database
func (repo *MysqlCourseRepository) DeleteCourse(courseID int) error {
	_, err := repo.DB.Exec("DELETE FROM courses WHERE id = ?", courseID)
	if err != nil {
		return fmt.Errorf("could not delete course: %v", err)
	}

	return nil
}

// GetCourseByID retrieves a course from the MySQL database by ID
func (repo *MysqlCourseRepository) GetCourseByID(courseID int) (model.Course, error) {
	var course model.Course

	err := repo.DB.QueryRow(
		"SELECT id, name, description, credits, instructor, schedule, capacity, available_seats FROM courses WHERE id = ?",
		courseID).Scan(&course.ID, &course.Name, &course.Description, &course.Credits, &course.Instructor,
		&course.Schedule, &course.Capacity, &course.AvailableSeats)
	if err != nil {
		if err == sql.ErrNoRows {
			return course, errors.New("no course found with the given ID")
		}
		return course, errors.New("error occurred while getting course data: " + err.Error())
	}
	return course, nil
}

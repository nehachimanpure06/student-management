package mysqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"student-management/pkg/model"
)

// MysqlStudentRepository implements StudentRepository using mysql database
type MysqlStudentRepository struct {
	DB *sql.DB
}

// NewMysqlStudentRepository creates a new MysqlStudentRepository instance
func NewMysqlStudentRepository(db *sql.DB) (*MysqlStudentRepository, error) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone VARCHAR(20),
    date_of_birth DATE,
    enrollment_date DATE,
    status ENUM('Active', 'Graduated', 'Dropped') NOT NULL)`)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating table students : %v", err.Error())
	}
	return &MysqlStudentRepository{DB: db}, nil
}

// GetAllStudents retrieves all students from the mysql database
func (repo *MysqlStudentRepository) GetAllStudents() ([]model.Student, error) {
	var students []model.Student
	rows, err := repo.DB.Query("SELECT id, first_name, last_name, email, phone, date_of_birth, enrollment_date, status FROM students")
	if err != nil {
		return nil, errors.New("error occured while getting student data : " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Phone, &student.DateOfBirth, &student.EnrollmentDate, &student.Status)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	return students, nil
}

// AddStudent adds a new student to the mysql database
func (repo *MysqlStudentRepository) AddStudent(student model.Student) (int, error) {
	query := `INSERT INTO students (first_name, last_name, email, phone, date_of_birth, enrollment_date, status)
	VALUES (?, ?, ?, ?, ?, ?, ?)`

	result, err := repo.DB.Exec(query, student.FirstName, student.LastName, student.Email, student.Phone, student.DateOfBirth, student.EnrollmentDate, student.Status)
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

// UpdateStudent updates an existing student in the mysql database
func (repo *MysqlStudentRepository) UpdateStudent(studentID int, student model.Student) error {
	query := `UPDATE students SET first_name = ?, last_name = ?, email = ?, phone = ?, date_of_birth = ?, enrollment_date = ?, status = ? WHERE id = ?`

	_, err := repo.DB.Exec(query, student.FirstName, student.LastName, student.Email, student.Phone, student.DateOfBirth, student.EnrollmentDate, student.Status, studentID)
	if err != nil {
		return fmt.Errorf("could not update student: %v", err)
	}

	return nil
}

// DeleteStudent deletes a student from the mysql database
func (repo *MysqlStudentRepository) DeleteStudent(studentID int) error {
	_, err := repo.DB.Exec("DELETE FROM students WHERE id = ?", studentID)
	if err != nil {
		return fmt.Errorf("could not delete student: %v", err)
	}

	return nil
}

// GetStudentByID retrieves a student from the mysql database by ID
func (repo *MysqlStudentRepository) GetStudentByID(studentId int) (model.Student, error) {
	var student model.Student

	err := repo.DB.QueryRow(
		"SELECT id, first_name, last_name, email, phone, date_of_birth, enrollment_date, status FROM students WHERE id = ?",
		studentId).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Phone,
		&student.DateOfBirth, &student.EnrollmentDate, &student.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return student, errors.New("no student found with the given ID")
		}
		return student, errors.New("error occurred while getting student data: " + err.Error())
	}
	return student, nil
}

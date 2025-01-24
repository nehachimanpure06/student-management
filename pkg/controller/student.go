package controller

import (
	"encoding/json"
	"strconv"
	"student-management/pkg/model/payload"
	"student-management/pkg/model/response"
	"student-management/pkg/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StudentController struct {
	studentService service.StudentService
}

// NewStudentController creates a new student controller
func NewStudentController(stdService service.StudentService) *StudentController {
	return &StudentController{
		studentService: stdService,
	}
}

func (s *StudentController) AddStudent(ctx *gin.Context) {
	var createStudentRequest payload.StudentRequest

	// Decode JSON request body into createStudentRequest struct
	err := json.NewDecoder(ctx.Request.Body).Decode(&createStudentRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "unable to decode the request data : "+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(createStudentRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "Invalid request body : "+err.Error())
		return
	}

	studentID, err := s.studentService.AddStudent(payload.ToStudentModel(createStudentRequest))
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to add student : "+err.Error())
		return
	}

	response.SuccessJSONResponse(ctx, response.IDResponse{ID: studentID})
}

func (s *StudentController) GetStudentByID(ctx *gin.Context) {
	studentId := ctx.Param("id")
	ID, err := strconv.Atoi(studentId)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid student id : "+err.Error())
		return
	}

	studentDetails, err := s.studentService.GetStudentByID(ID)
	if err != nil {
		response.InternalServerErrorJSON(ctx, "error occured while getting student details : "+err.Error())
		return
	}
	response.SuccessJSONResponse(ctx, response.ToStudentResponse(studentDetails))
}

func (s *StudentController) GetAllStudents(ctx *gin.Context) {
	students, err := s.studentService.GetStudents()
	if err != nil {
		response.InternalServerErrorJSON(ctx, err.Error())
		return
	}
	response.SuccessJSONResponse(ctx, response.ToStudentListResponse(students))
}

func (s *StudentController) UpdateStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")
	ID, err := strconv.Atoi(studentId)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid student id : "+err.Error())
		return
	}

	var updateStudentRequest payload.StudentRequest

	// Decode JSON request body into createStudentRequest struct
	err = json.NewDecoder(ctx.Request.Body).Decode(&updateStudentRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "unable to decode the request data : "+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(updateStudentRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "Invalid request body : "+err.Error())
		return
	}

	err = s.studentService.UpdateStudent(ID, payload.ToStudentModel(updateStudentRequest))
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to update user : "+err.Error())
		return
	}

	response.SuccessJSON(ctx, "student details updated successfully")

}

func (s *StudentController) DeleteStudent(ctx *gin.Context) {
	studentId := ctx.Param("id")
	ID, err := strconv.Atoi(studentId)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid student id : "+err.Error())
	}

	err = s.studentService.DeleteStudent(ID)
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to delete student details : "+err.Error())
		return
	}

	response.SuccessJSON(ctx, "student details deleted successfully")
}

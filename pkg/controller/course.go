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

type CourseController struct {
	courseService service.CourseService
}

// NewCourseController creates a new course controller
func NewCourseController(cService service.CourseService) *CourseController {
	return &CourseController{
		courseService: cService,
	}
}

func (s *CourseController) AddCourse(ctx *gin.Context) {
	var createCourseRequest payload.CourseRequest

	// Decode JSON request body into createCourseRequest struct
	err := json.NewDecoder(ctx.Request.Body).Decode(&createCourseRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "unable to decode the request data : "+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(createCourseRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "Invalid request body : "+err.Error())
		return
	}

	courseID, err := s.courseService.AddCourse(payload.ToCourseModel(createCourseRequest))
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to add course : "+err.Error())
		return
	}

	response.SuccessJSONResponse(ctx, response.IDResponse{ID: courseID})
}

func (s *CourseController) GetCourseByID(ctx *gin.Context) {
	courseID := ctx.Param("id")
	ID, err := strconv.Atoi(courseID)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid course id : "+err.Error())
		return
	}

	courseDetails, err := s.courseService.GetCourseByID(ID)
	if err != nil {
		response.InternalServerErrorJSON(ctx, "error occured while getting course details : "+err.Error())
		return
	}
	response.SuccessJSONResponse(ctx, response.ToCourseResponse(courseDetails))
}

func (s *CourseController) GetAllCourses(ctx *gin.Context) {
	courses, err := s.courseService.GetCourses()
	if err != nil {
		response.InternalServerErrorJSON(ctx, err.Error())
		return
	}
	response.SuccessJSONResponse(ctx, response.ToCourseListResponse(courses))
}

func (s *CourseController) UpdateCourse(ctx *gin.Context) {
	courseID := ctx.Param("id")
	ID, err := strconv.Atoi(courseID)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid course id : "+err.Error())
		return
	}

	var updateCourseRequest payload.CourseRequest

	// Decode JSON request body into createCourseRequest struct
	err = json.NewDecoder(ctx.Request.Body).Decode(&updateCourseRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "unable to decode the request data : "+err.Error())
		return
	}

	validator := validator.New()
	err = validator.Struct(updateCourseRequest)
	if err != nil {
		response.BadRequestJSON(ctx, "Invalid request body : "+err.Error())
		return
	}

	err = s.courseService.UpdateCourse(ID, payload.ToCourseModel(updateCourseRequest))
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to update user : "+err.Error())
		return
	}

	response.SuccessJSON(ctx, "course details updated successfully")

}

func (s *CourseController) DeleteCourse(ctx *gin.Context) {
	courseID := ctx.Param("id")
	ID, err := strconv.Atoi(courseID)
	if err != nil {
		response.BadRequestJSON(ctx, "invalid course id : "+err.Error())
	}

	err = s.courseService.DeleteCourse(ID)
	if err != nil {
		response.InternalServerErrorJSON(ctx, "Failed to delete course details : "+err.Error())
		return
	}

	response.SuccessJSON(ctx, "course details deleted successfully")
}

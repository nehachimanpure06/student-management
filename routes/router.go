package routes

import (
	"student-management/pkg/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, studentCtrl *controller.StudentController, courseCtrl *controller.CourseController) {
	students := r.Group("students")
	students.POST("/", studentCtrl.AddStudent)
	students.GET("/", studentCtrl.GetAllStudents)
	students.GET("/:id", studentCtrl.GetStudentByID)
	students.PUT("/:id", studentCtrl.UpdateStudent)
	students.DELETE("/:id", studentCtrl.DeleteStudent)

	courses := r.Group("courses")
	courses.POST("/", courseCtrl.AddCourse)
	courses.GET("/", courseCtrl.GetAllCourses)
	courses.GET("/:id", courseCtrl.GetCourseByID)
	courses.PUT("/:id", courseCtrl.UpdateCourse)
	courses.DELETE("/:id", courseCtrl.DeleteCourse)
}

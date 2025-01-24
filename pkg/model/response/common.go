package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	// StatusCode int    `json:"status_code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type IDResponse struct {
	ID int `json:"id"`
}

func SuccessJSONResponse(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusOK, object)
}

func SuccessJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Response{
		Status:  "OK",
		Message: message,
	})
}

func BadRequestJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusBadRequest, Response{
		Status:  "Bad request",
		Message: message,
	})
}

func NotFoundJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusNotFound, Response{
		Status:  "Not Found",
		Message: message,
	})
}

func InternalServerErrorJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusInternalServerError, Response{
		Status:  "Internal Server Error",
		Message: message,
	})
}

func ConflictErrorJSON(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusConflict, Response{
		Status:  "Conflict Error",
		Message: message,
	})
}

package response

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type successResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, &errorResponse{
		Message: message,
		Success: false,
	})
}

func NewSuccessResponse(c *gin.Context, statusCode int, data interface{}) {

	c.JSON(statusCode, &successResponse{
		Message: "Success",
		Success: true,
		Data:    data,
	})
}

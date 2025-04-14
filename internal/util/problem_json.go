package util

import (
	"github.com/free5gc/openapi/models"
	"github.com/gin-gonic/gin"
)

func GinProblemJson(c *gin.Context, problemDetails *models.ProblemDetails) {
	c.Writer.Header().Set("Content-Type", "application/problem+json")
	c.JSON(int(problemDetails.Status), problemDetails)
}

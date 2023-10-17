package utility

import (
	"github.com/gin-gonic/gin"
	"log"
)

func HttpInternalErrorResponse(c *gin.Context, msg string, err error) {
	log.Println(err)
	c.JSON(500, gin.H{
		"status":  "fail",
		"message": msg})
}

func HttpDataNotFound(c *gin.Context, msg string, err error) {
	c.JSON(404, gin.H{
		"status":  "fail",
		"message": msg,
		"data": gin.H{
			"error": err.Error(),
		},
	})
}

func HttpBadRequest(c *gin.Context, msg string) {
	c.JSON(400, gin.H{
		"status":  "fail",
		"message": msg,
	})
}

func HttpSuccessResponse(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"status":  "success",
		"message": msg,
		"data":    data,
	})
}

func HttpForbiddenResponse(c *gin.Context, msg string, err error) {
	c.JSON(403, gin.H{
		"status":  "fail",
		"message": msg,
		"data": gin.H{
			"err": err,
		},
	})
}

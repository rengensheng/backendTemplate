package utils

import (
	"github.com/rengensheng/backend/app/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResultError(code int, message interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": message,
		"type":    "fail",
	})
}

func ResultSuccess(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESSED,
		"message": "ok",
		"result":  data,
		"type":    "success",
	})
}

func ResultSuccessList(data interface{}, total int64, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    constant.SUCCESSED,
		"message": "ok",
		"result": gin.H{
			"items": data,
			"total": total,
		},
		"type": "success",
	})
}

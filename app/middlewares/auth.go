package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/utils"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if strings.HasPrefix(context.Request.URL.String(), "/ws") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/login") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/generator/database") ||
			strings.HasPrefix(context.Request.URL.String(), "/upload/") ||
			strings.HasPrefix(context.Request.URL.String(), "/api/public/") {
			context.Next()
			return
		}
		claims, err := utils.GetCurrentUserClaims(context)
		if err != nil {
			context.AbortWithStatusJSON(200, gin.H{
				"code":    constant.NO_ACCESS,
				"message": "登陆过期，请重新登陆",
			})
			return
		}
		context.Set("claims", claims)
		context.Next()
	}
}

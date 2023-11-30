package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rengensheng/backend/app/constant"
	"github.com/rengensheng/backend/app/utils"
	"strings"
)

var WhiteList = []string{
	"/api/user/login",
	"/api/upload/",
	"/api/public/",
	"/api/websocket"}

func isWhiteList(url string) bool {
	for _, v := range WhiteList {
		if strings.HasPrefix(url, v) {
			return true
		}
	}
	return false
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if isWhiteList(context.Request.URL.String()) {
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
		context.Set("username", claims.Username)
		context.Set("claims", claims)
		context.Next()
	}
}

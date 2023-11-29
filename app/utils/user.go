package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyahm/golog"
	"github.com/rengensheng/backend/app/models"
)

type CustomClaims struct {
	ID         string
	Username   string
	Permission []string
	Roles      []string
	jwt.StandardClaims
}

func GetCurrentUserClaims(c *gin.Context) (*CustomClaims, error) {
	cookie, err := c.Request.Cookie("token")
	token := ""
	if err != nil {
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := ParseToken(token)
	return userClaims, err
}

func GetCurrentUser(c *gin.Context) string {
	cookie, err := c.Request.Cookie("token")
	token := ""
	if err != nil {
		token = c.GetHeader("Authorization")
	} else {
		token = cookie.Value
	}
	userClaims, err := ParseToken(token)
	if err != nil {
		return "none"
	}
	return userClaims.Username
}

func GetPermCode(c *gin.Context) []string {
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		golog.Info(err.Error())
		return []string{}
	}
	userClaims, err := ParseToken(cookie.Value)

	if err != nil {
		golog.Info(err.Error())
		return []string{}
	}
	return userClaims.Permission
}

func GenerateToken(user models.User, permission, roleValues []string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3600000 * time.Second)
	golog.Info("permission", permission)
	issuer := user.Account
	claims := CustomClaims{
		ID:         user.Id,
		Username:   user.Account,
		Permission: permission,
		Roles:      roleValues,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    issuer,
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("golang"))
	return token, err
}

func ParseToken(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("golang"), nil
	})
	if err != nil {
		return &CustomClaims{}, err
	}
	if jwtToken != nil {
		if claims, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			return claims, nil
		}
	}
	return &CustomClaims{}, nil
}

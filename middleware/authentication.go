package middleware

import (
	"Valter/utility"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"strings"
)

type UserClaims struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorize := c.Request.Header.Get("Authorization")
		if !strings.HasPrefix(authorize, "Bearer ") {
			c.Abort()
			utility.HttpForbiddenResponse(c, "wrong header value", errors.New("wrong header"))
			log.Println("====================================")
			return
		}
		token := authorize[7:]
		claims := utility.UserClaims{}
		jwtKey := os.Getenv("SECRET_KEY")
		if _, err := utility.DecodeToken(token, &claims, jwtKey); err != nil {
			c.Abort()
			utility.HttpForbiddenResponse(c, "unauthorized", err)
			return
		}
		c.Set("user", claims)
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkc918/chat/response"
	"os"
)

func JWTTokenValid() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 验证token
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithResponse(c, response.ErrAccountAuthorizeFailed)
			c.Abort()
		}
		// 解析token
		parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		})
		if err != nil || !parse.Valid {
			response.FailWithResponse(c, response.ErrAccountAuthorizeFailed)
			c.Abort()
		}

		if claims, ok := parse.Claims.(jwt.MapClaims); ok {
			c.Set("user", claims)
		}
	}
}

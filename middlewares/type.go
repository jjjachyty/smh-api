package middlewares

import (
	"smh-api/middlewares/jwt"

	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("type", 0)

		if token := c.Request.Header.Get("authorization"); token != "" {
			j := jwt.NewJWT()
			// parseToken 解析token包含的信息

			if claims, err := j.ParseToken(token); err == nil && claims.UserID != 1 {
				c.Set("type", 1)
			}
		}
	}
}

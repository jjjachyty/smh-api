package middlewares

import (
	"fmt"
	"net/http"
	"smh-api/middlewares/jwt"

	"github.com/gin-gonic/gin"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("platform=", c.Request.Header.Get("platform"))
		if "ios" == c.Request.Header.Get("platform") {

			c.Set("type", 0)

			if token := c.Request.Header.Get("authorization"); token != "" {
				j := jwt.NewJWT()
				fmt.Println("authorization=", token)
				// parseToken 解析token包含的信息
				if claims, err := j.ParseToken(token); err == nil {
					fmt.Println("user=", claims.UserID)

					if claims.UserID != 1 {
						c.Set("type", 1)
					}
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{
						"State":   false,
						"Message": "授权已过期",
					})
					c.Abort()
					return
				}
			}
		}
	}
}

package middleware

import (
	"net/http"
	"strings"
	"week04/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请先登陆"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求头格式不对"})
			c.Abort()
			return
		}
		token := parts[1]
		//校验token
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户校验失败"})
		}
		idFloat := claims["id"].(float64)
		userId := uint(idFloat)
		c.Set("userId", userId)
		c.Set("username", claims["username"])
		return

	}
}

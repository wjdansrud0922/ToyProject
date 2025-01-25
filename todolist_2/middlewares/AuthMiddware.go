package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "토큰 인증 실패",
			})
			return
		}

		strtoken := strings.TrimPrefix(token, "Bearer ")

		tokenParse, err := jwt.Parse(strtoken, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("KEY")), nil
		})

		if err != nil || !tokenParse.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "토큰이 유효하질 않습니다",
			})
			return
		}

		if claim, ok := tokenParse.Claims.(jwt.RegisteredClaims); ok {
			username := claim.Subject
			c.Set("username", username)
			c.JSON(http.StatusOK, gin.H{"username": username})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "클레임 확인 실패"})
		}
		c.Next()
	}
}

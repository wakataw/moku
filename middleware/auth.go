package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wakataw/moku/pkg"
	"net/http"
	"strconv"
	"strings"
)

func ExtractToken(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	return tokenString
}

func AuthRequiredMiddleware(tokenManager *pkg.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := ExtractToken(c)
		token, err := tokenManager.Validate(tokenString)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userId, err := strconv.Atoi(claims["sub"].(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userId", userId)
		c.Set("roles", claims["roles"])
		c.Next()
	}
}

func isAdmin(roles any) bool {
	for _, v := range roles.([]interface{}) {
		if v.(string) == "admin" {
			return true
		}
	}

	return false
}

func AdminRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")

		if exists {
			if !isAdmin(roles) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		c.Next()
	}
}

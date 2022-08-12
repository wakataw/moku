package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/wakataw/moku/pkg"
	"log"
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

		roles := claims["roles"]
		log.Println(roles)

		c.Set("userId", userId)
		c.Set("roles", roles)
		c.Next()
	}
}

func isAdmin(roles []string) bool {
	log.Println(roles)
	for _, v := range roles {
		if v == "admin" {
			return true
		}
	}

	return false
}

func AdminRequiredMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, _ := c.Get("roles")

		if !isAdmin(roles) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}

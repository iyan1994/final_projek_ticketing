package middleware

import (
	"final_projek_ticketing/controller"
	"final_projek_ticketing/model"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminAuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Parse token

		claims, err := controller.ParseTokenMapClaims(tokenString)
		if err != nil {
			c.JSON(
				http.StatusUnauthorized,
				model.NewFailedResponse(fmt.Sprintf("Invalid or expired token: %s", err.Error())),
			)
			c.Abort()
			return
		}

		// Cek klaim token
		username := claims["username"].(string)
		idRole := int(claims["id_role"].(float64))
		title := claims["title"].(string)

		//otorisasi admin
		if idRole != 1 {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to create a user"})
			c.Abort()
			return
		}

		// Validasi token di database
		var user model.User
		if err := db.Where("username = ? AND token = ?", username, tokenString).First(&user).Error; err != nil {
			c.JSON(
				http.StatusUnauthorized,
				model.NewFailedResponse(fmt.Sprintf("Token not found in database: %s", err.Error())),
			)
			c.Abort()
			return
		}

		// Validasi token expired di database
		if user.ExpiredToken.Time.Before(time.Now()) {
			c.JSON(
				http.StatusUnauthorized,
				model.NewFailedResponse(fmt.Sprintf("Token expired in database: %s", err.Error())),
			)
			c.Abort()
			return
		}

		// Simpan klaim ke context
		c.Set("username", username)
		c.Set("id_role", idRole)
		c.Set("title", title)

		c.Next()
	}
}

package middlewares

import (
	"net/http"
	"zakki-store/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		claims := &config.JWTClaim{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, Token expired!"})
				c.Abort()
				return
			default:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func JWTMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		claims := &config.JWTAdmin{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY_Admin, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			case jwt.ValidationErrorExpired:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized, Token expired!"})
				c.Abort()
				return
			default:
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				c.Abort()
				return
			}
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

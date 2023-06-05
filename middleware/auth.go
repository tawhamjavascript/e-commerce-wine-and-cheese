package middleware

import (
	"e-commerce/config"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(c *fiber.Ctx) error {
    // Get the token string from the Authorization header
    authHeader := c.Get("Authorization")
    if authHeader == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Error to login user",
		})
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")

    // Parse and verify the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Check the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }

        // Return the secret key
        return []byte(config.Configs.SecretPassword), nil
    })
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Error to login user",
		})
    }

    // Get the claims from the token
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Error to login user",
		})
    }

    // Check the issuer, subject, and expiration time
    expirationTime := int64(claims["exp"].(float64))
    if time.Now().Unix() > expirationTime {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Token expired",
		})
    }

    // Set the user ID in the context for use in the protected route
    c.Locals("id", claims["sub"].(string))

    // Call the next middleware or route handler
    return c.Next()
}
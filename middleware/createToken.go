package middleware

import (
	"time"

	"e-commerce/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateToken(c *fiber.Ctx) error{
    // Define the issuer, subject, and expiration time
    issuer := "e-commerce of manga"
    expirationTime := time.Now().Add(24 * time.Hour).Unix()
    id := c.Locals("id").(primitive.ObjectID)

    // Create the token object with the specified claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
        "iss": issuer,
        "sub": id.Hex(),
        "exp": expirationTime,
    })
    // Sign the token with a secret key
    secretKey := []byte(config.Configs.SecretPassword)
    tokenString, err := token.SignedString(secretKey)
    if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error to login user",
		})
       
    }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User logged",
		"token": tokenString,
	})
}
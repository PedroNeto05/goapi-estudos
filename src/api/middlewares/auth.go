package middlewares

import (
	"errors"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func UserAuthenticated(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Voce precisa estar logado")
	}

	tokenString := strings.Split(authHeader, " ")[1]

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Token de autenticação invalido")
	}

	token, err := verifyToken(tokenString)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.Status(fiber.StatusInternalServerError).SendString("Token de autenticação invalido")
	}

	c.Locals("userID", claims["userID"])
	c.Locals("email", claims["email"])

	return c.Next()
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		secretKey := os.Getenv("JWT_SECRET_KEY")

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("método de assinatura inesperado")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token de autenticação invalido")
	}

	return token, nil
}

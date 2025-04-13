package middlewares

import (
	"strings"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/types"
)

type AuthMiddleware struct{}

type AuthMiddlewareInterface interface {
	BearerTokenExtraction(c fiber.Ctx) error
}

func (am *AuthMiddleware) BearerTokenExtraction(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	// Cek apakah header kosong
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  "Waduh, gak ketemu authorization header-nya mas! 😡",
		})
	}

	// Cek apakah formatnya bener "Bearer <token>"
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  "Hadeh, format authorization header-nya salah mas! 😡",
		})
	}
	tokenString := parts[1] // ambil token di bagian kedua format "Bearer <token>"

	// Decode token tanpa validasi
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  "Waduh, token-nya salah mas! 😡",
		})
	}

	// Ambil claims dari token alias decode semua isi token
	decodedToken, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  "Waduh, token-nya salah mas! 😡",
		})
	}

	// Simpan atribut yang dibutuhkan berdasarkan hasil decode token ke context
	email, ok := decodedToken["email"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(&types.CommonHandlerType{
			Response: false,
			Message:  "Waduh, token-nya salah mas! 😡",
		})
	}
	c.Locals("email", email)	

	// Lanjutkan ke handler berikutnya
	return c.Next()
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}
package middleware

import (
	"internship-mini-project/internal/exception"
	"internship-mini-project/internal/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewAuth(logger *logrus.Logger, Config *viper.Viper) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorization := c.Get("Authorization")

		bearerToken := strings.Split(authorization, " ")
		if bearerToken[0] != "Bearer" {
			logger.Error("invalid authorization header format")
			return fiber.ErrUnauthorized
		}

		token, err := jwt.Parse(bearerToken[1], func(t *jwt.Token) (any, error) {
			return []byte(Config.GetString("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			logger.WithError(err).Error("user unauthorized")
			return exception.ErrUserUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.WithError(err).Error("invalid token claims")
			return exception.ErrUserUnauthorized
		}

		auth := &model.Auth{
			ID:      uint(claims["id"].(float64)),
			Email:   claims["email"].(string),
			Nama:    claims["nama"].(string),
			NoTelp:  claims["no_telp"].(string),
			IsAdmin: claims["is_admin"].(bool),
		}

		c.Locals("auth", auth)

		return c.Next()
	}
}

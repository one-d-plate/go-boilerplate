package middleware

import (
	"crypto/rsa"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/one-d-plate/go-boilerplate.git/src/handler"
	"github.com/one-d-plate/go-boilerplate.git/src/pkg"
)

func GrantTokenVerify(pub *rsa.PublicKey) fiber.Handler {
	return func(c *fiber.Ctx) error {

		claims := jwt.MapClaims{}
		token := c.Get("Authorization")

		if token == "" {
			pkg.LogError("invalid grant token", nil)
			return handler.HandleFiberError(c, fiber.ErrUnauthorized)
		}

		jwt, err := pkg.ParseJwtToken(token, pub, claims)

		if err != nil {
			pkg.LogError("failed to parse JWT token", err)
			return handler.HandleFiberError(c, fiber.ErrUnauthorized)
		}

		expiredTime, err := jwt.GetExpirationTime()
		if err != nil {
			pkg.LogError("failed to parse JWT token", err)
			return handler.HandleFiberError(c, fiber.ErrUnauthorized)
		}

		if expiredTime.Before(time.Now()) {
			pkg.LogInfo("grant token is expired")
			return handler.HandleFiberError(c, fiber.ErrUnauthorized)
		}

		apiKey, ok := claims["api_key"].(string)
		if !ok {
			return handler.HandleFiberError(c, fiber.ErrUnauthorized)
		}

		c.Locals("clientId", apiKey)
		return c.Next()
	}
}

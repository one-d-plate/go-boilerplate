package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/one-d-plate/go-boilerplate.git/src/config"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

func RouteRegistry(app *fiber.App, db *bun.DB, rds *redis.Client) error {
	conf := config.NewConfig()
	conf.LoadFromEnv()

	return nil
}

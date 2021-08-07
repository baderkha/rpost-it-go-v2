package http

import "github.com/gofiber/fiber/v2"

func FiberRouteV2(fiberHandler func(*fiber.Ctx)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fiberHandler(c)
		return nil
	}
}

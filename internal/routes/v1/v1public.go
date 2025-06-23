package routes_api

import "github.com/gofiber/fiber/v2"

func PublicRouteV1(r fiber.Router) {
	routeV1 := r.Group("/v1/public")

	// Define all v1 public routes here
	routeV1.Get("/helo", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"msg": "hi public v1",
		})
	})
}

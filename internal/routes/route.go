package routes

import (
	"TrainTracking/internal/config"
	routes_api "TrainTracking/internal/routes/v1"
	"github.com/gofiber/fiber/v2"
)

func Register(r fiber.Router) {
	DB := config.GetDBConnection()
	api := r.Group("/api")

	routes_api.RouteV1(api, DB)
	routes_api.PublicRouteV1(api)
}

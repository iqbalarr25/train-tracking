package routes

import (
	routes_api "TrainTracking/internal/routes/v1"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(r fiber.Router, DB *gorm.DB) {
	api := r.Group("/api")

	routes_api.RouteV1(api, DB)
	routes_api.PublicRouteV1(api)
}

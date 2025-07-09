package routes_api

import (
	"TrainTracking/internal/config"
	"TrainTracking/internal/features/handler"
	"TrainTracking/internal/features/repository"
	"TrainTracking/internal/features/service"
	"TrainTracking/internal/routes/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RouteV1(r fiber.Router, DB *gorm.DB) {

	config.InitAWS()
	routeV1 := r.Group("/v1")
	auth := routeV1.Group("/auth")
	user := routeV1.Group("/users", middleware.CheckAuthDashboard)
	train := routeV1.Group("/trains")
	route := routeV1.Group("/routes")

	// Define all v1 routes here
	routeV1.Get("/helo", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"msg": "hi",
		})
	})

	userRepo := repository.NewUserRepository(DB)
	trainRepo := repository.NewTrainRepository(DB)
	routeRepo := repository.NewRouteRepository(DB)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	trainService := service.NewTrainService(trainRepo)

	routeService := service.NewRouteService(routeRepo)

	authHandler := handler.NewAuthHandler(authService)
	{
		auth.Post("/login", authHandler.Login)
	}

	userHandler := handler.NewUserHandler(userService)
	{
		user.Get("", userHandler.GetUserPagination)
		user.Post("", userHandler.CreateUser)
		user.Put("/:id", userHandler.UpdateUser)
		user.Delete("/:id", userHandler.DeleteUser)
	}

	trainHandler := handler.NewTrainHandler(trainService)
	{
		train.Get("", trainHandler.GetTrainPagination)
	}

	routeHandler := handler.NewRouteHandler(routeService)
	{
		route.Get("/:id", routeHandler.GetRouteById)
	}
}

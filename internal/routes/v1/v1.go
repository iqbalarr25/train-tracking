package routes_api

import (
	"TrainTracking/internal/config"
	"TrainTracking/internal/features/handler"
	"TrainTracking/internal/features/repository"
	"TrainTracking/internal/features/service"
	"TrainTracking/internal/routes/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/gofiber/websocket/v2"
)

func RouteV1(r fiber.Router, DB *gorm.DB) {

	config.InitAWS()
	routeV1 := r.Group("/v1")
	auth := routeV1.Group("/auth", middleware.CheckAuthDashboard)
	user := routeV1.Group("/users", middleware.CheckAuthDashboard)
	train := routeV1.Group("/trains", middleware.CheckAuthDashboard)
	route := routeV1.Group("/routes", middleware.CheckAuthDashboard)
	station := routeV1.Group("/stations", middleware.CheckAuthDashboard)
	ws := routeV1.Group("/ws")

	// Define all v1 routes here
	routeV1.Get("/helo", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(map[string]string{
			"msg": "hi",
		})
	})

	ws.Use(func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	userRepo := repository.NewUserRepository(DB)
	trainRepo := repository.NewTrainRepository(DB)
	routeRepo := repository.NewRouteRepository(DB)
	trackRepo := repository.NewTrackRepository(DB)
	stationRepo := repository.NewStationRepository(DB)

	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)

	trainService := service.NewTrainService(trainRepo)

	routeService := service.NewRouteService(routeRepo)

	trackService := service.NewTrackService(trackRepo)

	stationService := service.NewStationService(stationRepo)

	authHandler := handler.NewAuthHandler(authService)
	{
		auth.Post("/login", authHandler.Login)
		auth.Post("/register", authHandler.Register)
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

		ws.Get("/train/:id/positions", websocket.New(func(conn *websocket.Conn) {
			err := trainHandler.GetTrainPosition(conn)
			if err != nil {
				return
			}
		}))
	}

	routeHandler := handler.NewRouteHandler(routeService)
	{
		route.Get("/:id", routeHandler.GetRouteDetailByRouteId)
		trackHandler := handler.NewTrackHandler(trackService)
		{
			route.Get("/:id/tracks", trackHandler.GetTracksByRouteId)
		}
	}

	stationHandler := handler.NewStationHandler(stationService)
	{
		station.Get("/", stationHandler.GetStationPagination)
	}
}

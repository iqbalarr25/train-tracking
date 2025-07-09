package server

import (
	"TrainTracking/internal/scheduler"
	"TrainTracking/pkg/helper"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"os"

	"TrainTracking/internal/config"
	"TrainTracking/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var ctx = context.Background()

func StartApp() error {
	config.InitDatabase()
	config.InitCache(ctx)

	app := fiber.New()

	app.Use(logger.New())
	helper.StartTrainStatusScheduler(config.GetDBConnection())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))
	routes.Register(app)

	if err := scheduler.Register(); err != nil {
		helper.Exception(err)
		return err
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9000"
	}

	fmt.Println("Server running on port " + port)
	return app.Listen(":" + port)
}

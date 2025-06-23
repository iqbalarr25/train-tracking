package middleware

import (
	"TrainTracking/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
)

func CheckAuthDashboard(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return helper.GenerateResponse(ctx, http.StatusUnauthorized, fiber.ErrUnauthorized.Error(), nil, false)
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return helper.GenerateResponse(ctx, http.StatusUnauthorized, fiber.ErrUnauthorized.Error(), nil, false)
	}

	token, err := helper.VerifyToken(tokenString)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusUnauthorized, fiber.ErrUnauthorized.Error(), nil, false)
	}

	if !token.Valid {
		return helper.GenerateResponse(ctx, http.StatusUnauthorized, fiber.ErrUnauthorized.Error(), nil, false)
	}

	ctx.Locals("user", token)
	return ctx.Next()
}

package handler

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type (
	AuthHandlerInterface interface {
		Login(ctx *fiber.Ctx) error
	}

	AuthHandler struct {
		SVC service.AuthServiceInterface
	}
)

func NewAuthHandler(svc service.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{
		SVC: svc,
	}
}

func (h *AuthHandler) Login(ctx *fiber.Ctx) error {
	var request model.LoginAdminRequest
	validate := validator.New()
	if err := ctx.BodyParser(&request); err != nil {
		return helper.GenerateResponse(ctx, fiber.StatusBadRequest, err.Error(), nil, false)
	}

	if err := validate.Struct(&request); err != nil {
		return helper.GenerateResponse(ctx, fiber.StatusBadRequest, err.Error(), nil, false)
	}

	authResponse, statusCode, err := h.SVC.LoginAdmin(&request)
	if err != nil {
		return helper.GenerateResponse(ctx, statusCode, err.Error(), nil, false)
	}

	return helper.GenerateResponse(ctx, fiber.StatusOK, "", authResponse, true)
}

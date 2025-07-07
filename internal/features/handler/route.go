package handler

import (
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	RouteHandlerInterface interface {
		GetRouteById(ctx *fiber.Ctx) error
	}
	RouteHandler struct {
		SVC service.RouteServiceInterface
	}
)

func NewRouteHandler(svc service.RouteServiceInterface) RouteHandlerInterface {
	return &RouteHandler{
		SVC: svc,
	}
}

func (h *RouteHandler) GetRouteById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	resp, err := h.SVC.GetRouteById(id)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}

	return helper.GenerateResponse(ctx, http.StatusOK, helper.SuccessRetrievedDataMessage, resp, true)
}

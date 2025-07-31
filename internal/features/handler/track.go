package handler

import (
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type (
	TrackHandlerInterface interface {
		GetTracksByRouteId(ctx *fiber.Ctx) error
	}
	TrackHandler struct {
		SVC service.TrackServiceInterface
	}
)

func NewTrackHandler(svc service.TrackServiceInterface) TrackHandlerInterface {
	return &TrackHandler{
		SVC: svc,
	}
}

func (h *TrackHandler) GetTracksByRouteId(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	resp, err := h.SVC.GetTracksByRouteId(id)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}

	return helper.GenerateResponse(ctx, http.StatusOK, helper.SuccessRetrievedDataMessage, resp, true)
}

package handler

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
)

type (
	StationHandlerInterface interface {
		GetStationPagination(ctx *fiber.Ctx) error
	}
	StationHandler struct {
		SVC service.StationServiceInterface
	}
)

func NewStationHandler(svc service.StationServiceInterface) StationHandlerInterface {
	return &StationHandler{
		SVC: svc,
	}
}

func (h *StationHandler) GetStationPagination(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	pageSize := ctx.QueryInt("page_size", 5)
	search := ctx.Query("search")
	sort := ctx.Query("sort")
	filter := ctx.Query("filter")
	request := model.RequestPagination{
		Page:     page,
		PageSize: pageSize,
		Search:   search,
		Sort:     sort,
		Filter:   filter,
	}

	res, count, err := h.SVC.GetPagination(request)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}
	resp := helper.Response{
		Code:      http.StatusOK,
		Data:      res,
		Page:      int64(page),
		PageSize:  int64(pageSize),
		Total:     count,
		TotalPage: int64(math.Ceil(float64(count) / float64(pageSize))),
		Message:   helper.SuccessRetrievedDataMessage,
		Success:   true,
	}
	return helper.GenerateResponsePagination(ctx, resp)
}

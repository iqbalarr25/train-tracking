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
	TrainHandlerInterface interface {
		GetTrainPagination(ctx *fiber.Ctx) error
	}
	TrainHandler struct {
		SVC service.TrainServiceInterface
	}
)

func NewTrainHandler(svc service.TrainServiceInterface) TrainHandlerInterface {
	return &TrainHandler{
		SVC: svc,
	}
}

func (h *TrainHandler) GetTrainPagination(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	pageSize := ctx.QueryInt("page_size", 5)
	search := ctx.Query("search")
	sort := ctx.Query("sort")
	request := model.RequestPagination{
		Page:     page,
		PageSize: pageSize,
		Sort:     sort,
		Search:   search,
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

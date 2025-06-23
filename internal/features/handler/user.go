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
	UserHandlerInterface interface {
		CreateUser(ctx *fiber.Ctx) error
		UpdateUser(ctx *fiber.Ctx) error
		DeleteUser(ctx *fiber.Ctx) error
		GetUserPagination(ctx *fiber.Ctx) error
	}
	UserHandler struct {
		SVC service.UserServiceInterface
	}
)

func NewUserHandler(svc service.UserServiceInterface) UserHandlerInterface {
	return &UserHandler{
		SVC: svc,
	}
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var req model.AddUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return helper.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), nil, false)
	}
	_, err := h.SVC.CreateUser(req)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}
	return helper.GenerateResponse(ctx, http.StatusCreated, helper.SuccessCreatedDataMessage, nil, true)
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req model.UpdateUserRequest
	id, _ := ctx.ParamsInt("id")
	if err := ctx.BodyParser(&req); err != nil {
		return helper.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), nil, false)
	}
	err := h.SVC.UpdateUser(req, id)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}
	return helper.GenerateResponse(ctx, http.StatusCreated, helper.SuccessUpdatedDataMessage, nil, true)
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, _ := ctx.ParamsInt("id")
	err := h.SVC.DeleteUser(id)
	if err != nil {
		return helper.GenerateResponse(ctx, http.StatusInternalServerError, err.Error(), nil, false)
	}
	return helper.GenerateResponse(ctx, http.StatusOK, helper.SuccessDeletedDataMessage, nil, true)
}

func (h *UserHandler) GetUserPagination(ctx *fiber.Ctx) error {
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

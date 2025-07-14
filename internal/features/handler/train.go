package handler

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"github.com/gofiber/fiber/v2"
	"math"
	"net/http"
	"time"

	"github.com/gofiber/websocket/v2"
)

type (
	TrainHandlerInterface interface {
		GetTrainPagination(ctx *fiber.Ctx) error
		GetTrainPosition(conn *websocket.Conn) error
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
	filter := ctx.Query("filter")
	departStationId := ctx.Query("depart_station_id")
	arriveStationId := ctx.Query("arrive_station_id")
	request := model.RequestPaginationTrain{
		Page:            page,
		PageSize:        pageSize,
		Search:          search,
		Sort:            sort,
		Filter:          filter,
		DepartStationId: departStationId,
		ArriveStationId: arriveStationId,
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

func (h *TrainHandler) GetTrainPosition(conn *websocket.Conn) error {
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
		}
	}(conn)

	id := conn.Params("id")

	for {
		resp, err := h.SVC.GetTrainPosition(id)

		err = conn.WriteJSON(resp)
		if err != nil {
			helper.Exception(err)
			return nil
		}

		time.Sleep(5 * time.Second)
	}
}

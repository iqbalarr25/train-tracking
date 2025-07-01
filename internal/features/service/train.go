package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	TrainServiceInterface interface {
		GetPagination(req model.RequestPagination) (res []model.TrainListResponse, count int64, err error)
	}

	TrainService struct {
		Repo repository.TrainRepositoryInterface
	}
)

func NewTrainService(repo repository.TrainRepositoryInterface) TrainServiceInterface {
	return &TrainService{
		Repo: repo,
	}
}

func (s *TrainService) GetPagination(req model.RequestPagination) (res []model.TrainListResponse, count int64, err error) {
	var trains []model.Train
	count, err = s.Repo.GetTrainsPagination(req, &trains)

	// mapping ke response
	res = make([]model.TrainListResponse, len(trains))
	for i, t := range trains {
		res[i] = model.TrainListResponse{
			ID:         t.ID,
			Name:       t.Name,
			TrainCode:  t.TrainCode,
			Status:     t.Status,
			TrainClass: t.TrainClass,
			TrainRouteResponse: model.TrainListRouteResponse{
				ID:         t.Route.ID,
				DepartTime: t.Route.DepartTime,
				ArriveTime: t.Route.ArriveTime,
				DepartStation: model.TrainListStationResponse{
					ID:  t.Route.DepartStation.ID,
					Ref: t.Route.DepartStation.Ref,
				},
				ArriveStation: model.TrainListStationResponse{
					ID:  t.Route.ArriveStation.ID,
					Ref: t.Route.ArriveStation.Ref,
				},
			},
		}
	}

	return
}

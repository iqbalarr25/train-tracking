package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	StationServiceInterface interface {
		GetPagination(req model.RequestPagination) (res []model.StationListResponse, count int64, err error)
	}

	StationService struct {
		Repo repository.StationRepositoryInterface
	}
)

func NewStationService(repo repository.StationRepositoryInterface) StationServiceInterface {
	return &StationService{
		Repo: repo,
	}
}

func (s *StationService) GetPagination(req model.RequestPagination) (res []model.StationListResponse, count int64, err error) {
	var stations []model.Station
	count, err = s.Repo.GetStationsPagination(req, &stations)

	// mapping ke response
	res = make([]model.StationListResponse, len(stations))
	for i, s := range stations {
		res[i] = model.StationListResponse{
			ID:   s.ID,
			Name: s.Name,
			Ref:  s.Ref,
		}
	}

	return
}

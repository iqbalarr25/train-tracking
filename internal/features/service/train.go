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
	res, count, err = s.Repo.GetTrainsPagination(req)
	return
}

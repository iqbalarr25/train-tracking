package service

import (
	"TrainTracking/internal/features/repository"
)

type (
	SchedulerServiceInterface interface {
		UpdateTrainStatuses() error
	}

	schedulerService struct {
		RepoTrain repository.TrainRepositoryInterface
	}
)

func NewSchedulerService(repoTrain repository.TrainRepositoryInterface) SchedulerServiceInterface {
	return &schedulerService{
		RepoTrain: repoTrain,
	}
}

func (s *schedulerService) UpdateTrainStatuses() error {
	s.RepoTrain.UpdateTrainStatuses()
	return nil
}

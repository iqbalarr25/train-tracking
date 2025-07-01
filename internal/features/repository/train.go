package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
)

type (
	TrainRepositoryInterface interface {
		GetTrainsPagination(req model.RequestPagination, res *[]model.Train) (count int64, err error)
	}

	TrainRepository struct {
		DB *gorm.DB
	}
)

func NewTrainRepository(db *gorm.DB) *TrainRepository {
	return &TrainRepository{
		DB: db,
	}
}

func (r *TrainRepository) GetTrainsPagination(req model.RequestPagination, trains *[]model.Train) (count int64, err error) {
	query := r.DB.Model(&model.Train{}).Where("deleted_at is null").
		Preload("Route.DepartStation").
		Preload("Route.ArriveStation").
		Preload("TrainClass")

	query.Count(&count)

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if req.Sort != "" {
		query = query.Order(req.Sort)
	}

	err = query.Find(&trains).Error
	if err != nil {
		return
	}

	return
}

package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
)

type (
	TrainRepositoryInterface interface {
		GetTrainsPagination(req model.RequestPagination) (res []model.TrainListResponse, count int64, err error)
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

func (r *TrainRepository) GetTrainsPagination(req model.RequestPagination) (res []model.TrainListResponse, count int64, err error) {
	var trains []model.Train

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

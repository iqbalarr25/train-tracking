package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
)

type (
	RouteRepositoryInterface interface {
		GetRouteDetailByRouteId(id string, res *[]model.RouteDetail) (err error)
	}

	RouteRepository struct {
		DB *gorm.DB
	}
)

func NewRouteRepository(db *gorm.DB) *RouteRepository {
	return &RouteRepository{
		DB: db,
	}
}

func (r *RouteRepository) GetRouteDetailByRouteId(id string, routeDetail *[]model.RouteDetail) (err error) {
	err = r.DB.Where("route_id = ?", id).
		Preload("Station").
		Find(&routeDetail).Error
	if err != nil {
		return err
	}

	return
}

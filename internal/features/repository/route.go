package repository

import (
	"TrainTracking/internal/features/model"
	"fmt"
	"gorm.io/gorm"
)

type (
	RouteRepositoryInterface interface {
		GetRouteById(id string, res *model.Route) (err error)
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

func (r *RouteRepository) GetRouteById(id string, route *model.Route) (err error) {
	fmt.Print(123)
	err = r.DB.Model(model.Route{}).Where("id = ?", id).
		Preload("RouteDetails.Tracks").
		Preload("RouteDetails.Station").
		Preload("Train").
		Preload("DepartStation").
		Preload("ArriveStation").
		First(&route).Error

	return
}

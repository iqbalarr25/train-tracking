package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
)

type (
	StationRepositoryInterface interface {
		GetStationsPagination(req model.RequestPagination, res *[]model.Station) (count int64, err error)
	}

	StationRepository struct {
		DB *gorm.DB
	}
)

func NewStationRepository(db *gorm.DB) *StationRepository {
	return &StationRepository{
		DB: db,
	}
}

func (r *StationRepository) GetStationsPagination(req model.RequestPagination, stations *[]model.Station) (count int64, err error) {
	query := r.DB.Model(&model.Station{})

	if req.Search != "" {
		//query = query.Where(
		//	"trains.name ILIKE ? OR train_classes.name ILIKE ?",
		//	"%"+req.Search+"%", "%"+req.Search+"%",
		//)
	}

	if req.Filter != "" {
		//query = query.Where(
		//	"trains.status = ?",
		//	req.Filter,
		//)
	}

	query.Count(&count)

	offset := (req.Page - 1) * req.PageSize
	query = query.Offset(offset).Limit(req.PageSize)

	if req.Sort != "" {
		query = query.Order(req.Sort)
	}

	err = query.Find(&stations).Error
	if err != nil {
		return
	}

	return
}

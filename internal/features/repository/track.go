package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
)

type (
	TrackRepositoryInterface interface {
		GetTracksByRouteId(id string, res *[]model.Track) (err error)
	}

	TrackRepository struct {
		DB *gorm.DB
	}
)

func NewTrackRepository(db *gorm.DB) *TrackRepository {
	return &TrackRepository{
		DB: db,
	}
}

func (r *TrackRepository) GetTracksByRouteId(id string, tracks *[]model.Track) (err error) {
	err = r.DB.Joins("JOIN route_details ON route_details.id = tracks.route_detail_id").
		Where("route_details.route_id = ?", id).
		Preload("Station").
		Find(tracks).Error
	if err != nil {
		return err
	}

	return
}

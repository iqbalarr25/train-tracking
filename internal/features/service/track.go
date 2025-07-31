package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	TrackServiceInterface interface {
		GetTracksByRouteId(id string) (resp []model.GetTrackResponse, err error)
	}

	TrackService struct {
		Repo repository.TrackRepositoryInterface
	}
)

func NewTrackService(repo repository.TrackRepositoryInterface) TrackServiceInterface {
	return &TrackService{
		Repo: repo,
	}
}

func (s *TrackService) GetTracksByRouteId(id string) (res []model.GetTrackResponse, err error) {
	var tracks []model.Track
	err = s.Repo.GetTracksByRouteId(id, &tracks)

	for _, track := range tracks {
		var station *model.GetRouteTrackStationResponse
		if track.Station != nil {
			station = &model.GetRouteTrackStationResponse{
				ID:   track.Station.ID,
				Name: track.Station.Name,
				Ref:  track.Station.Ref,
			}
		}
		res = append(res, model.GetTrackResponse{
			ID:        track.ID,
			Sequence:  track.Sequence,
			Cost:      track.Cost,
			MaxSpeed:  track.MaxSpeed,
			Latitude:  track.Latitude,
			Longitude: track.Longitude,
			Station:   station,
		})
	}

	return res, err
}

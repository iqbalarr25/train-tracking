package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	RouteServiceInterface interface {
		GetRouteById(id string) (resp model.GetRouteResponse, err error)
	}

	RouteService struct {
		Repo repository.RouteRepositoryInterface
	}
)

func NewRouteService(repo repository.RouteRepositoryInterface) RouteServiceInterface {
	return &RouteService{
		Repo: repo,
	}
}

func (s *RouteService) GetRouteById(id string) (res model.GetRouteResponse, err error) {
	var route model.Route
	err = s.Repo.GetRouteById(id, &route)

	detailResponses := make([]model.GetRouteDetailResponse, 0)

	for _, detail := range route.RouteDetails {
		trackResponses := make([]model.GetRouteTrackResponse, 0)

		for _, track := range detail.Tracks {
			trackResponses = append(trackResponses, model.GetRouteTrackResponse{
				ID:        track.ID,
				Sequence:  track.Sequence,
				Cost:      float32(track.Cost),
				MaxSpeed:  track.MaxSpeed,
				Latitude:  track.Latitude,
				Longitude: track.Longitude,
			})
		}

		detailResponses = append(detailResponses, model.GetRouteDetailResponse{
			ID:         detail.ID,
			From:       detail.Station.Name,
			To:         detail.Station.Name,
			ArriveTime: detail.ArriveTime,
			DepartTime: detail.DepartTime,
			Track:      trackResponses,
		})
	}

	result := model.GetRouteResponse{
		ID:         route.ID,
		From:       route.DepartStation.Name,
		To:         route.ArriveStation.Name,
		ArriveTime: &route.ArriveTime,
		DepartTime: &route.DepartTime,
		Detail:     detailResponses,
	}

	return result, err
}

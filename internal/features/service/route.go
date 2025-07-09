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

	for i := 0; i < len(route.RouteDetails)-1; i++ {
		detail := route.RouteDetails[i]
		next := route.RouteDetails[i+1]

		trackResponses := make([]model.GetRouteTrackResponse, 0)
		for _, track := range detail.Tracks {
			trackResponses = append(trackResponses, model.GetRouteTrackResponse{
				ID:        track.ID,
				Sequence:  track.Sequence,
				MaxSpeed:  track.MaxSpeed,
				Cost:      float32(track.Cost),
				Latitude:  track.Latitude,
				Longitude: track.Longitude,
				StationID: track.StationID,
			})
		}

		detailResponses = append(detailResponses, model.GetRouteDetailResponse{
			ID:         detail.ID,
			From:       detail.Station.Name,
			To:         next.Station.Name,
			DepartTime: detail.DepartTime,
			ArriveTime: next.ArriveTime,
			Track:      trackResponses,
		})
	}

	result := model.GetRouteResponse{
		ID:         route.ID,
		From:       route.DepartStation.Name,
		To:         route.ArriveStation.Name,
		DepartTime: &route.DepartTime,
		ArriveTime: &route.ArriveTime,
		Detail:     detailResponses,
	}

	return result, err
}

package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
)

type (
	RouteServiceInterface interface {
		GetRouteDetailByRouteId(id string) (resp []model.GetRouteDetailResponse, err error)
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

func (s *RouteService) GetRouteDetailByRouteId(id string) (res []model.GetRouteDetailResponse, err error) {
	var routeDetails []model.RouteDetail
	err = s.Repo.GetRouteDetailByRouteId(id, &routeDetails)

	detailResponses := make([]model.GetRouteDetailResponse, 0)

	for _, routeDetail := range routeDetails {
		detailResponses = append(detailResponses, model.GetRouteDetailResponse{
			ID: routeDetail.ID,
			Station: model.GetRouteTrackStationResponse{
				ID:   routeDetail.ID,
				Name: routeDetail.Station.Name,
				Ref:  routeDetail.Station.Ref,
			},
			DepartTime: routeDetail.DepartTime,
			ArriveTime: routeDetail.ArriveTime,
			Latitude:   routeDetail.Station.Lat,
			Longitude:  routeDetail.Station.Lon,
		})
	}

	return detailResponses, err
}

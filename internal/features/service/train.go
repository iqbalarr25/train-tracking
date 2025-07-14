package service

import (
	"TrainTracking/internal/features/model"
	"TrainTracking/internal/features/repository"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"math"
	"sort"
	"time"
)

type (
	TrainServiceInterface interface {
		GetPagination(req model.RequestPaginationTrain) (res []model.TrainListResponse, count int64, err error)
		GetTrainPosition(id string) (res model.GetTrainPositionResponse, err error)
	}

	TrainService struct {
		Repo repository.TrainRepositoryInterface
	}
)

func NewTrainService(repo repository.TrainRepositoryInterface) TrainServiceInterface {
	return &TrainService{
		Repo: repo,
	}
}

func (s *TrainService) GetPagination(req model.RequestPaginationTrain) (res []model.TrainListResponse, count int64, err error) {
	var trains []model.Train
	count, err = s.Repo.GetTrainsPagination(req, &trains)

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
					ID:   t.Route.DepartStation.ID,
					Name: t.Route.DepartStation.Name,
					Ref:  t.Route.DepartStation.Ref,
				},
				ArriveStation: model.TrainListStationResponse{
					ID:   t.Route.ArriveStation.ID,
					Name: t.Route.ArriveStation.Name,
					Ref:  t.Route.ArriveStation.Ref,
				},
			},
		}
	}

	return
}

func (s *TrainService) GetTrainPosition(id string) (res model.GetTrainPositionResponse, err error) {
	var train model.Train
	err = s.Repo.GetTrainPosition(id, &train)
	if err != nil {
		return res, err
	}

	if train.Status == "Inactive" {
		return model.GetTrainPositionResponse{
			Lat:      0,
			Lon:      0,
			Progress: 0,
			Status:   train.Status,
		}, nil
	}

	if len(train.Route.RouteDetails) == 0 {
		return res, errors.New("tidak ada data detail route untuk posisi kereta")
	}

	detail := train.Route.RouteDetails[0]
	if detail.ArriveTime == nil || detail.DepartTime == nil {
		return res, errors.New("data waktu keberangkatan/kedatangan tidak lengkap")
	}

	if len(detail.Tracks) < 2 {
		return res, errors.New("jumlah track tidak cukup untuk menghitung posisi")
	}

	sort.Slice(detail.Tracks, func(i, j int) bool {
		return detail.Tracks[i].Sequence < detail.Tracks[j].Sequence
	})

	var coords []model.TrainLatLon
	for _, t := range detail.Tracks {
		coords = append(coords, model.TrainLatLon{Lat: t.Latitude, Lon: t.Longitude})
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	dummyDate := "2000-01-01"
	depart, _ := time.Parse("2006-01-02 15:04:05", dummyDate+" "+detail.DepartTime.Format("15:04:05"))
	arrive, _ := time.Parse("2006-01-02 15:04:05", dummyDate+" "+detail.ArriveTime.Format("15:04:05"))
	nowParsed, _ := time.Parse("2006-01-02 15:04:05", dummyDate+" "+now.Format("15:04:05"))
	fmt.Println(depart, arrive, nowParsed)
	totalDuration := arrive.Sub(depart).Seconds()
	elapsed := nowParsed.Sub(depart).Seconds()
	progress := elapsed / totalDuration

	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}

	log.Infof("🚆 Progress KA: %.2f%% (elapsed %.0fs dari %.0fs)", progress*100, elapsed, totalDuration)

	// Hitung jarak antar koordinat
	totalDist := 0.0
	distances := make([]float64, len(coords)-1)
	for i := 1; i < len(coords); i++ {
		d := Haversine(coords[i-1], coords[i])
		distances[i-1] = d
		totalDist += d
	}

	// Cari posisi berdasarkan progress
	target := progress * totalDist
	traveled := 0.0
	for i := 0; i < len(distances); i++ {
		if traveled+distances[i] >= target {
			remain := target - traveled
			ratio := remain / distances[i]
			lat := interpolate(coords[i].Lat, coords[i+1].Lat, ratio)
			lon := interpolate(coords[i].Lon, coords[i+1].Lon, ratio)

			if progress == 1 {
				train.Status = "Boarding"
			}
			return model.GetTrainPositionResponse{
				Lat:      lat,
				Lon:      lon,
				Progress: progress,
				Status:   train.Status,
			}, nil
		}
		traveled += distances[i]
	}

	last := coords[len(coords)-1]
	return model.GetTrainPositionResponse{
		Lat:      last.Lat,
		Lon:      last.Lon,
		Progress: progress,
		Status:   train.Status,
	}, nil
}

func Haversine(a, b model.TrainLatLon) float64 {
	const R = 6371e3 // meters
	dLat := (b.Lat - a.Lat) * math.Pi / 180
	dLon := (b.Lon - a.Lon) * math.Pi / 180

	lat1 := a.Lat * math.Pi / 180
	lat2 := b.Lat * math.Pi / 180

	h := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*math.Sin(dLon/2)*math.Sin(dLon/2)

	return R * 2 * math.Atan2(math.Sqrt(h), math.Sqrt(1-h))
}

func interpolate(a, b, t float64) float64 {
	return a + (b-a)*t
}

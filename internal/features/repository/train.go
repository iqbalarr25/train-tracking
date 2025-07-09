package repository

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
	"log"
	"time"
)

type (
	TrainRepositoryInterface interface {
		GetTrainsPagination(req model.RequestPaginationTrain, res *[]model.Train) (count int64, err error)
		UpdateTrainStatuses()
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

func (r *TrainRepository) GetTrainsPagination(req model.RequestPaginationTrain, trains *[]model.Train) (count int64, err error) {
	query := r.DB.Model(&model.Train{}).
		Where("trains.deleted_at IS NULL").
		Joins("JOIN train_classes ON train_classes.id = trains.train_class_id").
		Joins("JOIN routes ON routes.train_id = trains.id").
		Joins("JOIN stations AS depart_station ON depart_station.id = routes.depart_station_id").
		Joins("JOIN stations AS arrive_station ON arrive_station.id = routes.arrive_station_id")

	if req.Search != "" {
		query = query.Where(
			"trains.name ILIKE ? OR train_classes.name ILIKE ?",
			"%"+req.Search+"%", "%"+req.Search+"%",
		)
	}

	if req.Filter != "" {
		query = query.Where(
			"trains.status = ?",
			req.Filter,
		)
	}

	if req.DepartStationId != "" {
		query = query.Where(
			"depart_station.id = ?",
			req.DepartStationId,
		)
	}

	if req.ArriveStationId != "" {
		query = query.Where(
			"arrive_station.id = ?",
			req.ArriveStationId,
		)
	}

	query = query.
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

	return
}

func (r *TrainRepository) UpdateTrainStatuses() {
	var routes []model.Route

	err := r.DB.Preload("Train").Find(&routes).Error
	if err != nil {
		log.Println("Gagal ambil rute:", err)
		return
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	dummyDate := "2000-01-01"
	nowDummy, _ := time.ParseInLocation("2006-01-02 15:04:05", dummyDate+" "+now.Format("15:04:05"), loc)

	for _, route := range routes {
		departStr := route.DepartTime.Format("15:04:05")
		arriveStr := route.ArriveTime.Format("15:04:05")

		depart, _ := time.ParseInLocation("2006-01-02 15:04:05", dummyDate+" "+departStr, loc)
		arrive, _ := time.ParseInLocation("2006-01-02 15:04:05", dummyDate+" "+arriveStr, loc)
		board := depart.Add(-5 * time.Minute)

		status := "Inactive"
		switch {
		case nowDummy.Before(board) || nowDummy.After(arrive):
			status = "Inactive"
		case nowDummy.After(board) && nowDummy.Before(depart):
			status = "Boarding"
		case nowDummy.After(depart) && nowDummy.Before(arrive):
			status = "Moving"
		}

		log.Printf(
			"\n🚂 KA %s\nNow        : %s\nBoarding   : %s\nDepart     : %s\nArrive     : %s\nCurrentStat: %s\n",
			route.Train.Name,
			nowDummy.Format("15:04:05"),
			board.Format("15:04:05"),
			depart.Format("15:04:05"),
			arrive.Format("15:04:05"),
			route.Train.Status,
		)

		if route.Train.Status != status {
			err := r.DB.Table("trains").
				Where("id = ?", route.TrainID).
				Update("status", status).Error
			if err != nil {
				log.Printf("❌ Gagal update status KA %s: %v", route.Train.Name, err)
			} else {
				log.Printf("✅ Update status KA %s → %s", route.Train.Name, status)
			}
		}
	}
}

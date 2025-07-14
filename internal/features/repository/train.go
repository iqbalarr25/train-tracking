package repository

import (
	"TrainTracking/internal/features/model"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type (
	TrainRepositoryInterface interface {
		GetTrainsPagination(req model.RequestPaginationTrain, res *[]model.Train) (count int64, err error)
		GetTrainPosition(id string, train *model.Train) (err error)
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

//func (r *TrainRepository) GetTrainPosition(id string, train *model.Train) (err error) {
//	err = r.DB.Preload("Route").Where("id = ?", id).First(train).Error
//	if err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			return fmt.Errorf("train dengan ID '%s' tidak ditemukan: %w", id, err)
//		}
//		return fmt.Errorf("gagal mengambil data train: %w", err)
//	}
//
//	if train.Route == nil {
//		return errors.New("route tidak ditemukan atau tidak valid pada train")
//	}
//
//	var trainSummary model.TrainPositionRouteDetailSummary
//	err = r.getCurrentRouteDetail(train.Route.ID.String(), &trainSummary)
//	if err != nil {
//		return err
//	}
//
//	// Ambil semua RouteDetails hingga sequence saat ini
//	var routeDetails []model.RouteDetail
//	err = r.DB.
//		Where("route_id = ? AND sequence <= ?", train.Route.ID, trainSummary.CurrentSequence).
//		Order("sequence").
//		Preload("Tracks").
//		Find(&routeDetails).Error
//	if err != nil {
//		return fmt.Errorf("gagal preload route details: %w", err)
//	}
//
//	// Inject arrive_time ke detail yang aktif
//	for i := range routeDetails {
//		if routeDetails[i].ID == trainSummary.ID {
//			routeDetails[i].ArriveTime = &trainSummary.ArriveTime
//		}
//	}
//
//	train.Route.RouteDetails = routeDetails
//
//	return nil
//}

func (r *TrainRepository) GetTrainPosition(id string, train *model.Train) (err error) {
	err = r.DB.Preload("Route").Where("id = ?", id).First(train).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("train dengan ID '%s' tidak ditemukan: %w", id, err)
		}
		return fmt.Errorf("gagal mengambil data train: %w", err)
	}

	if train.Route == nil {
		return errors.New("route tidak ditemukan atau tidak valid pada train")
	}

	var trainSummary model.TrainPositionRouteDetailSummary
	err = r.getCurrentRouteDetail(train.Route.ID.String(), &trainSummary)
	if err != nil {
		return err
	}

	var routeDetails []model.RouteDetail
	err = r.DB.Where("route_id = ? AND sequence >= ? AND sequence <= ?", train.Route.ID, trainSummary.CurrentSequence, trainSummary.NextSequence).
		Preload("Tracks").
		Find(&routeDetails).Error
	if err != nil {
		return err
	}

	train.Route.RouteDetails = routeDetails

	return nil
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
		var detail model.TrainPositionRouteDetailSummary
		err := r.getCurrentRouteDetail(route.ID.String(), &detail)
		if err != nil {
			log.Printf("❌ Gagal dapat posisi detail untuk KA %s: %v", route.Train.Name, err)
			continue
		}

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

func (r *TrainRepository) getCurrentRouteDetail(id string, currentRoute *model.TrainPositionRouteDetailSummary) (err error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return fmt.Errorf("gagal memuat lokasi Jakarta: %w", err)
	}
	now := time.Now().In(loc)
	nowStr := now.Format("15:04:05")

	//nowStr := "18:52:00"

	query := `SELECT
	  rd_current.id,
	  rd_current.sequence AS current_sequence,
	  rd_current.depart_time AS depart_time,
	  (
		SELECT rd_next.sequence
		FROM route_details AS rd_next
		WHERE rd_next.route_id = rd_current.route_id
		  AND rd_next.sequence > rd_current.sequence
		  AND (rd_next.arrive_time IS NOT NULL OR rd_next.depart_time IS NOT NULL)
		ORDER BY rd_next.sequence
		LIMIT 1
	  ) AS next_sequence,
	  (
		SELECT COALESCE(rd_next.arrive_time, rd_next.depart_time)
		FROM route_details AS rd_next
		WHERE rd_next.route_id = rd_current.route_id
		  AND rd_next.sequence > rd_current.sequence
		  AND (rd_next.arrive_time IS NOT NULL OR rd_next.depart_time IS NOT NULL)
		ORDER BY rd_next.sequence
		LIMIT 1
	  ) AS arrive_time
	FROM
	  route_details AS rd_current
	WHERE
	  rd_current.route_id = $1
	  AND rd_current.depart_time IS NOT NULL
	  AND TO_CHAR(rd_current.depart_time::timestamp, 'HH24:MI:SS') <= $2
	ORDER BY
	  rd_current.sequence DESC
	LIMIT 1`

	err = r.DB.Raw(query, id, nowStr).Scan(&currentRoute).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("tidak ada posisi train yang cocok ditemukan untuk waktu saat ini")
		}
		return fmt.Errorf("gagal mengeksekusi query posisi train: %w", err)
	}

	return nil
}

package helper

import (
	"TrainTracking/internal/features/model"
	"gorm.io/gorm"
	"log"
	"time"
)

func StartTrainStatusScheduler(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			updateTrainStatuses(db)
		}
	}()
}

func updateTrainStatuses(db *gorm.DB) {
	var routes []model.Route

	err := db.Preload("Train").Find(&routes).Error
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
			err := db.Table("trains").
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

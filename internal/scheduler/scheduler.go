package scheduler

import (
	"TrainTracking/internal/config"
	"TrainTracking/internal/features/repository"
	"TrainTracking/internal/features/service"
	"TrainTracking/pkg/helper"
	"log"
)

func Register() error {
	config.InitScheduler()

	DB := config.GetDBConnection()

	cronScheduler, err := config.GetScheduler()
	if err != nil {
		return err
	}

	cronService := service.NewSchedulerService(
		repository.NewTrainRepository(DB),
	)
	_, err = cronScheduler.Every(1).Minute().Do(func() {
		err := cronService.UpdateTrainStatuses()
		if err != nil {
			log.Printf("❌ UpdateTrainStatuses error: %v", err)
		}
	})
	if err != nil {
		helper.Exception(err)
		return err
	}

	cronScheduler.StartAsync()
	return nil
}

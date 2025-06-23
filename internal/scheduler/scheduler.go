package scheduler

import (
	"TrainTracking/internal/config"
	"TrainTracking/pkg/helper"
)

func Register() error {
	config.InitScheduler()

	//DB := config.GetDBConnection()

	cronScheduler, err := config.GetScheduler()
	if err != nil {
		return err
	}

	//cronService := service.NewSchedulerService(
	//repository.NewOrbcomRepository(DB),
	//)
	//_, err = cronScheduler.Every(1).Day().At(config.GetString("FETCH_SATELLITE_SCHEDULER_TIME", "00:00")).Do(
	//	cronService.FetchSatelliteLog,
	//)

	if err != nil {
		helper.Exception(err)
		return err
	}

	cronScheduler.StartAsync()
	return nil
}

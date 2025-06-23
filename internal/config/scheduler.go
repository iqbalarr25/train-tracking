package config

import (
	"errors"
	"github.com/go-co-op/gocron"
	"time"
)

var scheduler *gocron.Scheduler

func InitScheduler() {
	scheduler = gocron.NewScheduler(time.UTC)
}

func GetScheduler() (*gocron.Scheduler, error) {
	if scheduler == nil {
		return nil, errors.New("scheduler not initialized")
	}

	return scheduler, nil
}

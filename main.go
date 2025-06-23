package main

import (
	"TrainTracking/cmd"
	"TrainTracking/internal/config"
)

func main() {
	config.InitEnv()
	cmd.Execute()
}

package main

import (
	"seowl/logger"
)

const DebugMode bool = true

func main() {

	// Init logger
	logger.InitLogger(DebugMode)
	log := logger.GetLogger()
	log.Debug("Logger Initialsed Sucessfully...")
}

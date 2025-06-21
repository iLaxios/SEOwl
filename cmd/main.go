package main

import (
	"seowl/crawler"
	"seowl/logger"
	"seowl/util"
)

const DebugMode bool = true

func main() {

	// Init logger
	logger.InitLogger(DebugMode)
	log := logger.GetLogger()
	log.Debug("Logger Initialsed Sucessfully...")

	// get user input
	url := "https://google.com"
	if !util.IsValidURL(url) {
		log.Error("User passed a invalid url!")
		return
	}

	// crawler
	crawler := crawler.NewCrawler(url)
	log.Debug(crawler)

}

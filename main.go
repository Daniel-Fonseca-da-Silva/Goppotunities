package main

import (
	"github.com/Daniel-Fonseca-da-Silva/Goppotunities/config"
	"github.com/Daniel-Fonseca-da-Silva/Goppotunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize config
	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize router
	router.Initialize()
}

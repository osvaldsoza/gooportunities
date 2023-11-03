package main

import (
	"github.com/osvaldsoza/gooportunities/config"
	"github.com/osvaldsoza/gooportunities/router"
)

var (
	loggeer *config.Logger
)

func main() {
	loggeer = config.GetLogger("main")
	err := config.Init()

	if err != nil {
		loggeer.Errorf("config initialization error: %v", err)
		return
	}

	router.Initialize()
}

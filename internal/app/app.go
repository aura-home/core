package app

import (
	"github.com/aura-home/core/internal/config"
	"github.com/aura-home/core/internal/modules/logger"
)

type App struct {
	configPath string
}

func NewApp(configPath string) *App {
	return &App{
		configPath: configPath,
	}
}

func (a App) MostStart() {
	cfg, err := config.Load(a.configPath)
	if err != nil {
		panic(err)
	}

	// Creating modules
	log := logger.New(&cfg.Core.Logger)

	log.Info("Starting Aura Home Core..")
	panic("Not Implemented")
}

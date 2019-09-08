package app

import (
	"github.com/maxvoronov/otus-go-10-logger/app/config"
	"github.com/maxvoronov/otus-go-10-logger/app/logger"
)

type App struct {
	Config *config.Config
	Logger *logger.Logger
}

func NewApp() (*App, error) {
	conf, err := config.NewConfig()
	if err != nil {
		return nil, err
	}

	logs, err := logger.NewLogger(conf)
	if err != nil {
		return nil, err
	}

	app := &App{
		Config: conf,
		Logger: logs,
	}

	return app, nil
}

func (app *App) Start() {
	app.Logger.Infof("Start application in [%s] mode...", app.Config.AppMode)
	// Do something useful
	app.Logger.Info("Done!")
}

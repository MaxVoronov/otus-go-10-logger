package logger

import (
	"github.com/maxvoronov/otus-go-10-logger/app/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(conf *config.Config) (*Logger, error) {
	var zapLogger *zap.Logger
	var err error
	if conf.IsDevMode() {
		if zapLogger, err = zap.NewDevelopment(); err != nil {
			return nil, err
		}
	} else {
		if zapLogger, err = zap.NewProduction(); err != nil {
			return nil, err
		}
	}

	defer zapLogger.Sync()
	logger := &Logger{zapLogger.Sugar()}

	return logger, nil
}

package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func LogInit(level string) error {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	Log = logger.Sugar()
	return nil
}

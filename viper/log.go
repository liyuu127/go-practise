package viper

import "go.uber.org/zap"

var Sugar *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment(zap.AddCaller())
	defer logger.Sync() // flushes buffer, if any
	Sugar = logger.Sugar()
}

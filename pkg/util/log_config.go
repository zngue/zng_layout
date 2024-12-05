package util

import "github.com/zngue/zng_app/log"

func LogConfig() *log.Config {
	configDefault := log.WriterConfigDefault
	configDefault.Level = log.LevelInfo
	return configDefault
}

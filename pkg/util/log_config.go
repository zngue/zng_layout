package util

import (
	"github.com/zngue/zng_app/log"
	"io"
	"os"
)

var logMap = map[string]log.LevelType{
	"debug": log.LevelDebug,
	"info":  log.LevelInfo,
	"warn":  log.LevelWarn,
	"error": log.LevelError,
}

func LogConfig() *log.Config {
	configDefault := log.WriterConfigDefault
	level := os.Getenv("LOG_LEVEL")
	if logMap[level] > 0 {
		configDefault.Level = logMap[level]
	}
	configDefault.WriteSyncer = LogWriter()
	return configDefault
}
func LogWriter() io.Writer {
	var url = os.Getenv("LOG_URL")
	if url == "" {
		return nil
	}
	return log.NewLogSave(url)

}

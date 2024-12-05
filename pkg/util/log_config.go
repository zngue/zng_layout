package util

import (
	"github.com/zngue/zng_app/log"
	"io"
	"os"
)

func LogConfig() *log.Config {
	configDefault := log.WriterConfigDefault
	configDefault.Level = log.LevelDebug
	configDefault.WriteSyncer = LogWriter()
	return configDefault
}
func LogWriter() io.Writer {
	var url = os.Getenv("LOG_URL")
	if url == "" {
		url = "http://localhost:16666/v1/log/create"
	}
	return log.NewLogSave(url)

}

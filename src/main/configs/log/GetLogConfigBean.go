package main_configs_log

import (
	"log/slog"
	"os"
	"sync"
)

const _MSG_INITIALIZING_LOG_BEANS = "Initializing logs configuration beans"
const _MSG_LOG_BEANS_INITIATED = "Log configuration beans successfully initiated"

var once sync.Once
var logConfigBean *slog.Logger

func GetLogConfigBean() *slog.Logger {
	once.Do(func() {
		if logConfigBean == nil {
			logConfigBean = getLogConfig()
		}
	})
	return logConfigBean
}

func getLogConfig() *slog.Logger {
	slog.Info(_MSG_INITIALIZING_LOG_BEANS)
	logLevel := new(slog.LevelVar)
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	//Como é ponteiro, mudando o ponteiro muda todos mundo que usa ele tbm
	logLevel.Set(slog.LevelDebug)
	slog.SetDefault(l)

	// "!BADKEY": pois sempre tem que ser key/value
	slog.Info(_MSG_LOG_BEANS_INITIATED)

	return l
}

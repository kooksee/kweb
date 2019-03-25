package kweb

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func (t *app) InitLog() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	if !t.IsDebug {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	}

	t.Ip = IpAddress()
	if t.Ip == "" {
		panic("获取不到ip地址")
	}

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service_name", t.ServiceName).
		Str("service_ip", t.Ip).
		Bool("is_debug", t.IsDebug).
		Caller().
		Logger()
}

package kweb

import (
	"github.com/kooksee/kweb/internal/g"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func (t *app) InitLog() {
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"
	zerolog.SetGlobalLevel(g.If(t.IsDebug, zerolog.DebugLevel, zerolog.ErrorLevel).(zerolog.Level))

	t.Ip = g.IpAddress()
	g.AssertBool(t.Ip == "", "获取不到ip地址")

	log.Logger = log.
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		With().
		Str("service_name", t.ServiceName).
		Str("service_ip", t.Ip).
		Bool("is_debug", t.IsDebug).
		Caller().
		Logger()
}

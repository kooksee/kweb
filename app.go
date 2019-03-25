package kweb

import (
	"github.com/jmoiron/sqlx"
	"github.com/kooksee/go-assert"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type app struct {
	ServiceName string
	IsDebug     bool
	Ip          string
	dbs         map[string]*sqlx.DB
}

func (t *app) InitDb(cfg *DbConfig) {
	db, err := sqlx.Connect(cfg.Schema, cfg.DbUrl)
	assert.Err(err, "init db %s error")
	db.SetMaxIdleConns(10)
	t.dbs[cfg.Name] = db
}

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

var _app *app
var once sync.Once

func InitApp() *app {
	return GetApp()
}

func GetApp() *app {

	once.Do(func() {
		_app = &app{
			ServiceName: serviceName,
			IsDebug:     true,
		}
	})

	return _app
}

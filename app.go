package kweb

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

type app struct {
	ServiceName string
	IsDebug     bool
	Ip          string
	dbs         map[string]*sqlx.DB
}

func (t *app) Start() {

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

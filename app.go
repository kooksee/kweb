package kweb

import (
	"sync"
)

type app struct {
	Cfg *Config
}

func (t *app) InitLog() {

}

var _app *app
var once sync.Once

func DefaultApp() *app {

	once.Do(func() {
		_app = &app{}
	})
	return _app
}

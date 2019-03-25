package kweb

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type app struct {
	Cfg *Config
}

func (t *app) InitLog() {

}

func (t *app) InitConfig() {
	// 环境变量配置
	viper.SetEnvPrefix("kweb")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	// 配置文件
	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/kweb/")
	viper.AddConfigPath("$HOME/.kweb")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

var _app *app
var once sync.Once

func DefaultApp() *app {

	once.Do(func() {
		_app = &app{}
	})
	return _app
}

package kweb

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type KError struct {
	Code string
	Msg  string
}

type KErrors struct {
	data map[string]map[string][]string
}

func (t *KErrors) FromPath(cfg string) {
	_, err := toml.DecodeFile(cfg, &t.data)
	assertErr(err, "文件[%s]解析失败", cfg)
}

func (t *KErrors) Get(ns, name string, args ...interface{}) *KError {
	return &KError{
		Code: t.data[ns][name][0],
		Msg:  fmt.Sprintf(t.data[ns][name][1], args...),
	}
}

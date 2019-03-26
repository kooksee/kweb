package kweb

import (
	"github.com/BurntSushi/toml"
	"github.com/antonmedv/expr"
	"github.com/kooksee/kweb/internal/g"
	"github.com/kooksee/kweb/internal/validator"
	"io"
)

type KForm struct {
	Parser expr.Node
	Msg    string
}

type KForms struct {
	data map[string]map[string]*KForm
}

func (t *KForms) FromPath(cfg string) {
	var _dt map[string]map[string][]string

	_, err := toml.DecodeFile(cfg, &_dt)
	g.AssertErr(err, "文件[%s]解析失败", cfg)

	for k := range _dt {
		for j := range _dt[k] {
			p, err := expr.Parse(_dt[k][j][0], expr.Env(&validator.KValidator{}))
			g.AssertErr(err, "规则[%s]解析失败", _dt[k][j][0])
			t.data[k][j] = &KForm{Parser: p, Msg: _dt[k][j][1]}
		}
	}
}

func (t *KForms) Validator(form string, r io.Reader) string {
	var _dt map[string]interface{}
	g.AssertErr(g.Json.NewDecoder(r).Decode(&_dt), "json解析失败")

	for k, v := range t.data[form] {
		if _d, ok := _dt[k]; ok {
			if _s := validator.KValidatorOf(_d).Do(v.Parser); _s != "" {
				return _s
			}
		}
	}

	return ""
}

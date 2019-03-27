package kweb

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/antonmedv/expr"
	"github.com/kooksee/kweb/internal/g"
	"github.com/kooksee/kweb/internal/validator"
	"io"
	"strings"
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
			var _fn []expr.OptionFn

			if !strings.HasPrefix(j, "__") {
				_fn = append(_fn, expr.Env(&validator.KValidator{}))
			}

			p, err := expr.Parse(_dt[k][j][0], _fn...)
			g.AssertErr(err, "规则[%s]解析失败", _dt[k][j][0])
			t.data[k][j] = &KForm{Parser: p, Msg: _dt[k][j][1]}
		}
	}
}

func (t *KForms) Validator(form string, r io.Reader) string {
	var _dt map[string]interface{}
	g.AssertErr(g.Json.NewDecoder(r).Decode(&_dt), "json解析失败")

	for k, v := range t.data[form] {
		_s := ""
		_b := true

		if strings.HasPrefix(k, "__") {
			_b, _s = validator.KValidatorOf(_dt).Eval(v.Parser)
		}

		if _d, ok := _dt[k]; ok {
			_b, _s = validator.KValidatorOf(_d).Do(v.Parser)
		}

		if !_b {
			return g.If(_s == "", v.Msg, fmt.Sprintf(v.Msg+", Err:%s", _s)).(string)
		}

	}

	return ""
}

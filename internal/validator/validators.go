package validator

import (
	"github.com/antonmedv/expr"
	"github.com/kooksee/kweb/internal/g"
	"reflect"
)

func KValidatorOf(d interface{}) *KValidator {
	return &KValidator{data: reflect.ValueOf(d)}
}

type KValidator struct {
	data reflect.Value
	err  string
}

func (t *KValidator) do(node expr.Node, dt interface{}) (bool, string) {
	out, err := expr.Run(node, dt)
	g.AssertErr(err, "校验规则执行失败")

	ok, isBool := out.(bool)
	g.AssertBool(!isBool, "校验规则结果类型错误")

	return ok, t.err
}

func (t *KValidator) Do(node expr.Node) (bool, string) {
	return t.do(node, t)
}

func (t *KValidator) Eval(node expr.Node) (bool, string) {
	return t.do(node, t.data)
}

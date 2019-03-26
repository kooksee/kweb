package tests

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/antonmedv/expr"
	"github.com/kooksee/g/assert"
	"io/ioutil"
	"testing"
)

func TestName(t *testing.T) {
	dt, err := ioutil.ReadFile("errors.toml")
	assert.Err(err, "")

	var _d map[string]interface{}

	_md, err := toml.Decode(string(dt), &_d)
	assert.Err(err, "")
	assert.P(_md)
	assert.P(_d)
}

type wp struct {
	Data string
}

func (h *wp) Title(s string) bool {
	fmt.Println(s)
	fmt.Println(h.Data)
	return true
}

func TestName1(t *testing.T) {
	p, err := expr.Parse("!Title('sss')", expr.Env(&wp{}))
	assert.MustNotError(err)

	out, err := expr.Run(p, &wp{Data: "world"})
	assert.MustNotError(err)
	assert.P(out)
}

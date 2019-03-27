package tests

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/antonmedv/expr"
	"github.com/gin-gonic/gin"
	"github.com/kooksee/g/assert"
	"github.com/kooksee/kweb/internal/g"
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

func Test2(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	g.Assert(r.Run(":8888"))
}


//https://www.baidu.com/link?url=V8RO3dPz-yIEbjaomd69xA7Q1B3IbpNDB1pkP5IDEFZ392ihsiNkrXC8Dq-TgKcMeySMHYYq0fiETM-RCLD5pHuQHOKP1Ol4KK6XNTY0Lc3


//V8RO3dPz
// yIEbjaomd69xA7Q1B3IbpNDB1pkP5IDEFZ392ihsiNkrXC8Dq
// TgKcMeySMHYYq0fiETM
// RCLD5pHuQHOKP1Ol4KK6XNTY0Lc3
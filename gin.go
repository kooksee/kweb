package kweb

import (
	"github.com/gin-gonic/gin"
	"github.com/kooksee/kweb/internal/g"
)

func (t *app) app() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		dt, err := c.GetRawData()
		g.AssertErr(err, "")
		// check form
		// do sql
		// do return
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    dt,
		})
	})
	g.Assert(r.Run())
}

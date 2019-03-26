package kweb

import "github.com/gin-gonic/gin"

func (t *app) app() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		dt, err := c.GetRawData()
		assertErr(err, "")

		c.JSON(200, gin.H{
			"message": "pong",
			"data":    dt,
		})
	})
	assert(r.Run())
}

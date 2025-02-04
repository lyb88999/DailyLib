package main

import (
	"github.com/sohaha/zlsgo/znet"
	"net/http"
)

func main() {
	r := znet.New()

	// 设置为开发模式
	r.SetMode(znet.DebugMode)

	// 异常处理
	r.Use(znet.Recovery(func(c *znet.Context, err error) {
		e := err.Error()
		c.String(http.StatusInternalServerError, e)
	}))

	// 注册路由
	r.GET("/json", func(c *znet.Context) {
		c.JSON(http.StatusOK, znet.Data{"message": "Hello world!"})
	})

	r.GET("/", func(c *znet.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	znet.Run()
}

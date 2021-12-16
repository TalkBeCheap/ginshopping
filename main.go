package main

import (
	"ginshopping/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default使用了log中间价， 再调用new
	app := gin.Default()
	app.LoadHTMLGlob("template/**/**") // 挂载所有

	app.Static("/static", "static") // 挂载css文件夹,静态文件挂载
	//router.LoadHTMLFiles("index.html")
	app.GET("/", router.Hello)
	app.GET("/user", router.User)
	app.GET("/userinfo", router.UserDetails)
	app.GET("/param/:id", router.Params)
	app.GET("/query/", router.Query)
	app.GET("/queryarray", router.QueryArray)
	app.GET("/querymap", router.QueryMap)
	//router := gin.New()
	app.Run()
}

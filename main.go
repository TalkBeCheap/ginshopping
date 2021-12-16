package main

import (
	. "ginshopping/middleware"
	"ginshopping/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// Default使用了log中间价， 再调用new
	app := gin.Default()
	app.LoadHTMLGlob("template/**/**") // 挂载所有
	app.Use(TimeMiddleware,TimeMiddleware2,TimeMiddleware3)
	app.Static("/static", "static") // 挂载css文件夹,静态文件挂载
	//router.LoadHTMLFiles("index.html")
	app.GET("/", router.Hello)
	app.GET("/user", router.User)
	app.GET("/userinfo", router.UserDetails)
	app.GET("/param/:id", router.Params)
	app.GET("/query/", TimeConuntMiddware,router.Query)
	app.GET("/queryarray", router.QueryArray)
	app.GET("/querymap", router.QueryMap)
	app.GET("/form", router.PostFrom)
	app.POST("/formadd", router.PostFromAdd) // 表单提交POST请求
	app.POST("/FormToStructure", router.FormToStructure) // 表单提交POST请求
	app.GET("/jsonp", router.TransJsonp)
	app.GET("/purejson", router.PureJson)
	app.GET("/asciijson", router.AsciiJSON)
	app.GET("/yaml", router.Yaml)
	app.GET("/redirect", router.Redirect)
	//router := gin.New()
	app.Run()
}

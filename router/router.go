package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func Hello(ctx *gin.Context) {
	name := "youku"

	ctx.HTML(200, "index/index.html", name)
}

func User(ctx *gin.Context) {
	ctx.HTML(200, "user/user.html", nil)
}

func UserDetails(ctx *gin.Context) {
	userInfo := UserInfo{1, "showing", 18, "party"}
	ctx.HTML(http.StatusOK, "user/info.html", userInfo)
}

func Params(ctx *gin.Context) {
	// http://127.0.0.1:8080/param/0 param参数获取
	id := ctx.Param("id")
	ctx.String(200, "Hello, %s", id)
}

func Query(ctx *gin.Context) {
	//http://127.0.0.1:8080/query/?title=3 query参数获取
	query := ctx.Query("title")
	query2 := ctx.DefaultQuery("name", "jerky") //没有传则代表默认参数
	ctx.String(200, "title %s, %s", query, query2)
}

func QueryArray(ctx *gin.Context) {
	//http://127.0.0.1:8080/queryarray?id=12,123
	query := ctx.QueryArray("id")
	ctx.String(200, "title %s", query)
}

func QueryMap(ctx *gin.Context) {
	//http://127.0.0.1:8080/querymap?map[id]=3&map[age]=18
	query := ctx.QueryMap("map")
	ctx.String(200, "map %v", query)
}

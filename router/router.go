package router

import (
	"fmt"
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

func PostFrom(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "form/form.html", nil)

}

// POST请求参数获取
func PostFromAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	fmt.Printf("username: %v\npassword: %v\n", username, password)
	ctx.String(http.StatusOK, "添加成功")
}

// 接受表单提交的数据
type UserAbout struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

func FormToStructure(ctx *gin.Context) {
	var u UserAbout
	err := ctx.ShouldBind(&u) // 绑定到结构体
	fmt.Println(err)
	fmt.Println(u)
	ctx.String(http.StatusOK, "Hello !%s", u.Name)
}

func TransJsonp(ctx *gin.Context) {
	data := map[string]interface{}{
		"foo": "bar",
	}

	// /JSONP?callback=x
	// 将输出：x({\"foo\":\"bar\"})
	ctx.JSONP(http.StatusOK, data)
}

func PureJson(ctx *gin.Context) {
	ctx.PureJSON(200, gin.H{
		//返回{"html":"<b>Hello, world!</b>"}
		"html": "<b>Hello, world!</b>",
	})
}

func AsciiJSON(ctx *gin.Context) {
	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

	// 输出 : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
	ctx.AsciiJSON(http.StatusOK, data)
}

func Yaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
}

func Redirect(ctx *gin.Context) {
	ctx.Redirect(http.StatusFound, "/user")
}
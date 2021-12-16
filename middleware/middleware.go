package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)
//type HandlerFunc func(*Context)
func TimeMiddleware(ctx *gin.Context) {
	fmt.Println("这是自定义中间件1start")
	//if 3 < 4 {
	//	ctx.Abort() //补，不满足条件终止下一个中间件的流向
	//}
	ctx.Next()
	fmt.Println("这是自定义中间件1ed")

}

func TimeMiddleware2(ctx *gin.Context) {
	fmt.Println("这是自定义中间件2start")
	ctx.Next()
	fmt.Println("这是自定义中间件2end")
}

func TimeMiddleware3(ctx *gin.Context) {
	fmt.Println("这是自定义中间件3start")
	ctx.Next()
	fmt.Println("这是自定义中间件3end")
}

func TimeConuntMiddware(ctx *gin.Context) {
	start := time.Now()
	ctx.Next()
	duration := time.Since(start)
	fmt.Printf("time 耗时 %v", duration)
}

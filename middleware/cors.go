//package middleware
//
//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func CORSMiddleware() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
//		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")          //设置缓存时间
//		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")        //可以通过访问的方法  *表示所有
//		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")        //指定header带的信息
//		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true") //??????
//
//		if ctx.Request.Method == http.MethodOptions {
//			ctx.AbortWithStatus(200)
//		} else {
//			ctx.Next()
//		}
//
//	}
//}

package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(
		cors.Config{
			//AllowAllOrigins:  true,
			AllowOrigins:     []string{"*"}, // 等同于允许所有域名 #AllowAllOrigins:  true
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*","Authorization"},
			ExposeHeaders:    []string{"Content-Length", "text/plain", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
	)
}


package routes

import "github.com/gin-gonic/gin"

func IndexRoute(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Hello World!")
	})
}

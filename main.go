package main

import (
	"github.com/ericwangdq/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(200, "Index page")
	// })

	routes.ConfigRoute(router)
	routes.IndexRoute(router)
	router.Run(":9090")
}

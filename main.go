package main

import (
	"github.com/ericwangdq/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	routes.IndexRoute(router)
	routes.ConfigRoute(router)

	router.Run(":9090")
}

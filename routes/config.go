package routes

import (
	"fmt"
	"net/http"

	"github.com/ericwangdq/gin-gorm-rest/middlewares"
	"github.com/ericwangdq/gin-gorm-rest/utils"
	"github.com/gin-gonic/gin"
)

const RGM_CONFIG_PATH = "./config/config_rgp.json"
const RGM_USER_PATH = "./config/user.json"
const RGM_CREATE_USER_PATH = "./config/create.json"
const RGM_USER_ME_PATH = "./config/me.json"
const RGM_ROLES_PATH = "./config/roles.json"

func ConfigRoute(router *gin.Engine) {
	router.GET("/config.json", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGM_CONFIG_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

	router.POST("/auth/csrf-token", func(ctx *gin.Context) {

		token := map[string]interface{}{
			"token": "qm5jh9QFUSt7S46G",
		}

		ctx.AsciiJSON(http.StatusOK, token)
	})

	router.GET("/auth/user", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGM_USER_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

	api := router.Group("/api/")
	{
		api.GET("/user-configuration-plus/business/v1/users/create", func(ctx *gin.Context) {

			err, resp := utils.ReadFile(RGM_CREATE_USER_PATH)

			if err == nil {
				ctx.JSON(http.StatusOK, resp)
			} else {
				ctx.JSON(http.StatusInternalServerError, resp)
			}
		})

		api.GET("/user-configuration-plus/business/v1/users/me", func(ctx *gin.Context) {

			err, resp := utils.ReadFile(RGM_USER_ME_PATH)

			if err == nil {
				ctx.JSON(http.StatusOK, resp)
			} else {
				ctx.JSON(http.StatusInternalServerError, resp)
			}
		})

		api.GET("/user-configuration-plus/business/v1/users/roles", func(ctx *gin.Context) {

			err, resp := utils.ReadFile(RGM_ROLES_PATH)

			if err == nil {
				ctx.JSON(http.StatusOK, resp)
			} else {
				ctx.JSON(http.StatusInternalServerError, resp)
			}
		})
	}

	router.Any("/api/:upstream/*path", middlewares.Internal(), func(ctx *gin.Context) {
		upstream := ctx.Param("upstream")
		path := ctx.Param("path")
		fmt.Println("Upstream: " + upstream)
		fmt.Println("Path: " + path)
		if upstream == "/user-configuration-plus" {
			ctx.Next()
		} else {
			fmt.Println("No such API: " + path)
			msg := map[string]interface{}{
				"msg": " No such API",
			}
			ctx.JSON(http.StatusNotFound, msg)
		}
	})

	// router.GET("/api/*", func(ctx *gin.Context) {
	// 	fmt.Println("No such API: " + ctx.FullPath())
	// 	ctx.String(404, "No such API")
	// })
}

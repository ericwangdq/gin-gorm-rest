package routes

import (
	"net/http"

	"github.com/ericwangdq/gin-gorm-rest/utils"
	"github.com/gin-gonic/gin"
)

const RGP_CONFIG_PATH = "./config/config_rgp.json"
const RGP_USER_PATH = "./config/user.json"
const RGP_CREATE_USER_PATH = "./config/create.json"
const RGP_USER_ME_PATH = "./config/me.json"
const RGP_ROLES_PATH = "./config/roles.json"

func ConfigRoute(router *gin.Engine) {
	router.GET("/config.json", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGP_CONFIG_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
		// ctx.String(200, `<html>
		// 	<h1>
		// 	config.json
		// 	</h1>
		// </html>`)
	})

	router.POST("/auth/csrf-token", func(ctx *gin.Context) {

		token := map[string]interface{}{
			"token": "qm5jh9QFUSt7S46G",
		}

		ctx.AsciiJSON(http.StatusOK, token)
	})

	router.GET("/auth/user", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGP_USER_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

	router.GET("/api/user-configuration-plus/business/v1/users/create", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGP_CREATE_USER_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

	router.GET("/api/user-configuration-plus/business/v1/users/me", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGP_USER_ME_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

	router.GET("/api/user-configuration-plus/business/v1/users/roles", func(ctx *gin.Context) {

		err, resp := utils.ReadFile(RGP_ROLES_PATH)

		if err == nil {
			ctx.JSON(http.StatusOK, resp)
		} else {
			ctx.JSON(http.StatusInternalServerError, resp)
		}
	})

}
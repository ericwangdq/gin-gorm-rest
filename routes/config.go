package routes

import (
	"fmt"
	"net/http"
	"net/http/httputil"

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

	// api := router.Group("/api/")
	// {
	// 	api.GET("/user-configuration-plus/business/v1/users/create", func(ctx *gin.Context) {

	// 		err, resp := utils.ReadFile(RGM_CREATE_USER_PATH)

	// 		if err == nil {
	// 			ctx.JSON(http.StatusOK, resp)
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, resp)
	// 		}
	// 	})

	// 	api.GET("/user-configuration-plus/business/v1/users/me", func(ctx *gin.Context) {

	// 		err, resp := utils.ReadFile(RGM_USER_ME_PATH)

	// 		if err == nil {
	// 			ctx.JSON(http.StatusOK, resp)
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, resp)
	// 		}
	// 	})

	// 	api.GET("/user-configuration-plus/business/v1/users/roles", func(ctx *gin.Context) {

	// 		err, resp := utils.ReadFile(RGM_ROLES_PATH)

	// 		if err == nil {
	// 			ctx.JSON(http.StatusOK, resp)
	// 		} else {
	// 			ctx.JSON(http.StatusInternalServerError, resp)
	// 		}
	// 	})
	// }

	router.Any("/api/:upstream/*path", middlewares.Internal(), func(ctx *gin.Context) {
		upstream := ctx.Param("upstream")
		path := ctx.Param("path")
		path = fmt.Sprintf("/api/%s%s", upstream, path)
		fmt.Println("Upstream: " + upstream)
		fmt.Println("Path: " + path)

		host := "rgp.rgp-uat.e4844d9.kyma.ondemand.com"
		cookie := "gin=MTY3NzU5MDI4MHxOd3dBTkVzMVZVNU1XVm8zV2xCUVRGaEZRbE5EVFVWRFIwMVBXRTlJVERRMlRsbFZTVlpCTTB0TFFsZFZVRTAzU1U1WVRWWlpObEU9fK8WmfZbvhtdonHxVK5JlssgwCiGztGfFJUddHyGHMdy"
		proxy := &httputil.ReverseProxy{
			Director: func(req *http.Request) {
				req.URL.Scheme = "https"
				req.Host = "rgp.rgp-uat.e4844d9.kyma.ondemand.com"
				req.URL.Host = "rgp.rgp-uat.e4844d9.kyma.ondemand.com"
				req.URL.Path = path
				req.Header.Set("X-Forwarded-Host", host)
				req.Header.Set("Forwarded", fmt.Sprintf("host=%s", host))
				req.Header.Set("cookie", cookie)
				// req.Header.Set(consts.DOMAIN_HEADER, domain)
				// if !configData.AuthConfig.SkipBackend {

				// 	req.Header.Set("Authorization", Bearer+jwtToken)
				// }

				// fillin dynamic headers
				// fillInDynamicHeader(req, configData, claims)

				// correctAndLogRequest(req, traceId)
				fmt.Println("API call - Director: " + path)
			},

			ModifyResponse: func(rsp *http.Response) error {
				// if secWebSocketProtocols != "" {
				// 	rsp.Header.Set(SecWebSocketProtocol, secWebSocketProtocols)
				// }

				// log.Info().Msgf("Response from %s, status code %d, trace id: %s, headers %v", rsp.Request.URL.String(), rsp.StatusCode, traceId, rsp.Header)

				// // Clear session if backend service check token failed, so that user would be able to try to login again.
				// if rsp.StatusCode == http.StatusUnauthorized {
				// 	utils.CleanLoginInfo(c)
				// }
				fmt.Println("API call - ModifyResponse: " + path)
				return nil
			},
		}
		proxy.ServeHTTP(ctx.Writer, ctx.Request)

		// upstream == "/rgp-promotion-config" ||
		// if upstream == "/rgp-promotion-planning" {
		// 	// ctx.Next()

		// } else {

		// 	fmt.Println("No such API: " + path)
		// 	msg := map[string]interface{}{
		// 		"msg": " No such API",
		// 	}
		// 	ctx.JSON(http.StatusNotFound, msg)
		// }

	})

	// router.GET("/api/*", func(ctx *gin.Context) {
	// 	fmt.Println("No such API: " + ctx.FullPath())
	// 	ctx.String(404, "No such API")
	// })
}

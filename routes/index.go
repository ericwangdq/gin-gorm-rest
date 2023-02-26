package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoute(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		// render index.html
		// response, err := http.Get("https://cdn.industry.cloud.sap/rgp/rgp-shell-ui/ee57900/index.html")
		response, err := http.Get("http://127.0.0.1:3031/shell/index.html")
		if err != nil || response.StatusCode != http.StatusOK {
			ctx.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		// extraHeaders := map[string]string{
		// 	"Content-Disposition": `attachment; filename="gopher.png"`,
		// }

		extraHeaders := map[string]string{}

		ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		// download image
		// response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		// if err != nil || response.StatusCode != http.StatusOK {
		// 	ctx.Status(http.StatusServiceUnavailable)
		// 	return
		// }

		// reader := response.Body
		// contentLength := response.ContentLength
		// contentType := response.Header.Get("Content-Type")

		// extraHeaders := map[string]string{
		// 	"Content-Disposition": `attachment; filename="gopher.png"`,
		// }

		// ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

		// 	ctx.String(200, `<html>
		// 	<h1>
		// 		{{ .title }}
		// 	</h1>
		// </html>`)

		// 	ctx.HTML(200, `<html>
		// 	<h1>
		// 		{{ .title }}
		// 	</h1>
		// </html>`, gin.H{
		// 		"title": "RGM Suite",
		// 	})
	})

}

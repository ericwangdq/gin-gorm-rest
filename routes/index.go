package routes

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ericwangdq/gin-gorm-rest/utils"
	"github.com/gin-gonic/gin"
)

const FavoriteIcon = "data:image/vnd.microsoft.icon;base64,AAABAAEAEBAAAAEAIACwAQAAFgAAAIlQTkcNChoKAAAADUlIRFIAAAAQAAAAEAgDAAAAKC0PUwAAAAFzUkdCAdnJLH8AAAAJcEhZcwAACxMAAAsTAQCanBgAAACHUExURSIiIiIiIv6/APy9AP+/APy9AP+/AAAAABwcHOasAP+/AP2+AP+/AAAAACIiIv+/AP29AP+/APy9AP+/AP+/AP29AP+/APy9AP+/APu8AP+/APy9AP+/APu8AP+/APy9AP+/AP+/AP+/ACIiItukAB8fHz82Hv+/APy9AP+/AP6+APy9AP6/ADvzOYsAAAAtdFJOUw8EDL2+v58ACxf//94CKxXNmJl8gt65uliB0J9Ig927u7gIChAYPZuenZ6aBuT+UGgAAACJSURBVHicjY7bEgExEESDSDMxIkRcFpu16xL8//fJ7pI3VR5O1fRDnx4hBsORlHLcooCJAKZEpDtm4LkAG6KFtXZp7YoBwTCkHZDujlQxmlzO3Dtczugda+/9xvttdnx2dtmxV4VSqjj8dhxz/vuPUxlCKEN1To66dXx3m1S54Hq7xxgTj+cLeAOuNgzqSaOekwAAAABJRU5ErkJggg=="
const Title = "Revenue Growth Management Sales Planning"
const Description = ""

func IndexRoute(router *gin.Engine) {
	router.Use(func(ctx *gin.Context) {
		if ctx.Request.Method == http.MethodGet {
			CdnPrefix := "https://cdn.industry.cloud.sap/rgp/rgp-shell-ui/ee57900"
			err, config := utils.ReadFile(RGM_CONFIG_PATH)

			if err == nil {
				fmt.Printf("config: ", config)
				shellUi, ok := config["shell-ui"]
				if !ok {
					fmt.Printf("shell UI does not exist\n")
					return
				}
				shellUiMap := shellUi.(map[string]interface{})
				fmt.Printf("shell-ui config: ", shellUiMap["location"].(string))
				CdnPrefix = shellUiMap["location"].(string)
			} else {
				fmt.Printf("could not unmarshal config json: %s\n", err)
				return
			}

			// render index.html
			response, err := http.Get("https://cdn.industry.cloud.sap/rgp/rgp-shell-ui/ee57900/index.html")
			// response, err := http.Get("http://127.0.0.1:3031/shell/index.html")
			if err != nil || response.StatusCode != http.StatusOK {
				ctx.Status(http.StatusServiceUnavailable)
				return
			}

			defer response.Body.Close()
			bodyBytes, bodyErr := io.ReadAll(response.Body)
			if bodyErr != nil {
				fmt.Println(bodyErr)
			}
			bodyString := string(bodyBytes)
			replacer := strings.NewReplacer("{{.FavoriteIcon}}", FavoriteIcon, "{{.CdnPrefix}}", CdnPrefix, "{{.Title}}", Title, "{{.Description}}", Description)
			replacedBody := replacer.Replace(bodyString)

			contentType := response.Header.Get("Content-Type")
			ctx.Data(http.StatusOK, contentType, []byte(replacedBody))
		} else {
			ctx.Next()
		}
	})

	// router.GET("/", func(ctx *gin.Context) {
	// 	// download image
	// 	// response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
	// 	// if err != nil || response.StatusCode != http.StatusOK {
	// 	// 	ctx.Status(http.StatusServiceUnavailable)
	// 	// 	return
	// 	// }

	// 	// reader := response.Body
	// 	// contentLength := response.ContentLength
	// 	// contentType := response.Header.Get("Content-Type")

	// 	// extraHeaders := map[string]string{
	// 	// 	"Content-Disposition": `attachment; filename="gopher.png"`,
	// 	// }

	// 	// ctx.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	// 	// 	ctx.String(200, `<html>
	// 	// 	<h1>
	// 	// 		{{ .title }}
	// 	// 	</h1>
	// 	// </html>`)

	// 	// 	ctx.HTML(200, `<html>
	// 	// 	<h1>
	// 	// 		{{ .title }}
	// 	// 	</h1>
	// 	// </html>`, gin.H{
	// 	// 		"title": "RGM Suite",
	// 	// 	})
	// })

}

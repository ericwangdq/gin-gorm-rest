package main

import (
	"encoding/json"
	"fmt"

	"github.com/ericwangdq/gin-gorm-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// router.GET("/*", func(ctx *gin.Context) {
	// 	ctx.String(200, "Index page")
	// })

	jsonData := `
		{
			"intValue":1234,
			"boolValue":true,
			"stringValue":"hello!",
			"dateValue":"2022-03-02T09:10:00Z",
			"objectValue":{
				"arrayValue":[1,2,3,4]
			},
			"nullStringValue":null,
			"nullIntValue":null
		}
	`

	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	fmt.Printf("json map: %v\n", data)
	rawDateValue, ok := data["dateValue"]
	if !ok {
		fmt.Printf("dateValue does not exist\n")
		return
	}
	dateValue, ok := rawDateValue.(string)
	if !ok {
		fmt.Printf("dateValue is not a string\n")
		return
	}
	fmt.Printf("date value: %s\n", dateValue)

	fmt.Printf("objectValue: %#v\n", data["objectValue"])

	routes.ConfigRoute(router)
	routes.IndexRoute(router)
	router.Run(":9090")
}

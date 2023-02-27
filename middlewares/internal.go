package middlewares

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// Internal try to get consts.InternalAccessFF
// If FF not exist reject /internal/* call
// If FF exist and set false to current tenant reject /internal/* call
// If FF exist and set true to current tenant allow /internal/* call
func Internal() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !isInternalAPIs(ctx) {
			ctx.Next()
			return
		}
	}
}

func isInternalAPIs(ctx *gin.Context) bool {
	path := ctx.Param("path")

	internalAPIList := []string{"/internal"}
	for _, value := range internalAPIList {
		if strings.Index(path, value) == 0 {
			fmt.Println("Internal API: " + ctx.FullPath())
			return true
		}
	}
	return false
}

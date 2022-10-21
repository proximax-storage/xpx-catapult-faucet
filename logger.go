package Faucet

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proximax-storage/xpx-catapult-faucet/utils"
)

func ParseLogger(inner gin.HandlerFunc, name string) gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		defer utils.PanicCtrl()

		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		if raw != "" {
			path = path + "?" + raw
		}

		utils.Logger(1,
			"%s %s %s %s",
			ctx.Request.Method,
			path,
			name,
			time.Since(start),
		)
		inner(ctx)
	})
}

func SetCORS(active bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if active {
			if ctx.Request.Header.Get("Origin") != "" {
				ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
			} else {
				ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			}

			if ctx.Request.Header.Get("Access-Control-Request-Headers") != "" {
				ctx.Writer.Header().Set("Access-Control-Allow-Headers", ctx.Request.Header.Get("Access-Control-Request-Headers"))
			}

			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
			ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization, TokenRefresh, ExpiresIn")

			if ctx.Request.Method == "OPTIONS" {
				ctx.AbortWithStatus(204)
				return
			}
			ctx.Next()
		}
	}
}

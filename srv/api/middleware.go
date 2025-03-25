package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ranjbar-dev/gobit/config"
)

func (a *Api) registerMiddlewares() {

	router := a.hs.GetRouter()

	// register basic-auth middleware
	router.Use(basicAuthMiddleware)
}

func whitelistIPMiddleware(ctx *gin.Context) {

	if ctx.Request.URL.Path == "/server/ping" {

		ctx.Next()
		return
	}

	// abort request
	ctx.AbortWithStatus(403)
}

func basicAuthMiddleware(ctx *gin.Context) {

	if ctx.Request.URL.Path == "/server/ping" {

		ctx.Next()
		return
	}

	username, password, hasAuth := ctx.Request.BasicAuth()
	if !hasAuth || username != config.ApiAuthUsername() || password != config.ApiAuthPassword() {

		ctx.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		ctx.AbortWithStatus(401)
		return
	}

	ctx.Next()
}

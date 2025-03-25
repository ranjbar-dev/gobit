package api

import (
	docs "github.com/ranjbar-dev/gobit/.swagger/docs"
	"github.com/ranjbar-dev/gobit/config"
	apicontroller "github.com/ranjbar-dev/gobit/srv/api/controllers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Api) registerSwagger() {

	docs.SwaggerInfo.Title = config.SwaggerTitle()
	docs.SwaggerInfo.Description = "Responses indicate the success or failure of the request, with HTTP status codes of 200 for success, 400 for bad input, and 500 for errors."
	docs.SwaggerInfo.Version = config.SwaggerVersion()
	docs.SwaggerInfo.Host = config.SwaggerHost()
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{config.SwaggerProtocol()}

	router := a.hs.GetRouter()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func (a *Api) registerRoutes() {

	controller := apicontroller.NewController()

	// register middlewares //

	a.registerMiddlewares()

	// swagger //

	a.registerSwagger()

	// server //

	a.hs.RegisterGetRoute("/server/ping", controller.ServerPingPost)

	a.hs.RegisterPostRoute("/server/ping", controller.ServerPingPost)

	a.hs.RegisterGetRoute("/server/timestamp", controller.ServerTimestamp)

	a.hs.RegisterGetRoute("/server/memory-usage", controller.ServerMemoryUsage)

}

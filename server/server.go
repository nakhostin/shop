package server

import (
	"fmt"
	"micro_services/identity/api"
	"micro_services/identity/constant"
	"micro_services/identity/context"
	"micro_services/identity/server/config"
	"micro_services/identity/static"
	"micro_services/identity/utils/color"
	"micro_services/msbase/db/mongo"
	"micro_services/msbase/logger"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Run() {
	config.InitConfig()

	loadMiddlewares()
	echo.NotFoundHandler = static.NotFound

	fmt.Printf("%sBinary Version %s%s\n", color.Blue, constant.Version, color.Reset)
	fmt.Printf("%s%s %s%s\n", color.Blue, config.CFG.App, config.CFG.Version, color.Reset)
	// command.Init()

	//setup logger configs and files
	logger.Setup()

	//setting up database name
	mongo.SetDatabase(config.CFG.MongoDB)
	//create connection with mongo database , uri will going to get from Conf/app.yaml file
	mongo.Connect(config.CFG.MongoHost)

	api.Routes(constant.EC)

	constant.EC.Start(fmt.Sprintf(":%s", config.CFG.ListenerPort))
}

func loadMiddlewares() {
	middlewares := []echo.MiddlewareFunc{}

	middlewares = append(middlewares, context.ContextHandler)
	if config.CFG.Debug {
		middlewares = append(middlewares, middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: middleware.DefaultCORSConfig.AllowMethods,
		}))

		// middlewares = append(middlewares, middleware.Logger())
	} else {
		middlewares = append(middlewares, middleware.CORSWithConfig(middleware.CORSConfig{
			AllowMethods: config.CFG.AllowMethods,
		}))
	}

	constant.EC.Use(middlewares...)
}

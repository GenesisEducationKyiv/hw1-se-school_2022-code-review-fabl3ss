package main

import (
	"orders/src/config"
	"orders/src/pkg/application"
	"orders/src/pkg/delivery/http/handlers"
	"orders/src/pkg/delivery/http/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func StartApi() {
	router := createRouter()
	handlers := createHandlers()

	routes.InitRoutes(
		router,
		handlers,
	)

	startServer(router)
}

func createRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	return gin.Default()
}

func createHandlers() *handlers.Handlers {
	ordersUsecase := application.NewOrdersUsecase(
		os.Getenv(config.EnvDtmCoordinatorURL),
		os.Getenv(config.EnvCustomersServiceURL),
	)

	return &handlers.Handlers{
		OrdersHandler: handlers.NewOrdersHandler(ordersUsecase),
	}
}

func startServer(engine *gin.Engine) {
	err := engine.Run(os.Getenv(config.EnvServerURL))
	if err != nil {
		panic("unable to start http server")
	}
}

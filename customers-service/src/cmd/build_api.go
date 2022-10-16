package main

import (
	"customers/src/config"
	"customers/src/pkg/application"
	"customers/src/pkg/delivery/http/handlers"
	"customers/src/pkg/delivery/http/routes"
	"customers/src/pkg/persistence"
	"customers/src/platform"
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
	dbURL := os.Getenv(config.EnvMySqlURL)
	db := platform.NewMySqlDB(dbURL)

	customerRepo := persistence.NewCustomerRepository(db)
	customerUsecase := application.NewCustomerUsecase(customerRepo)
	customerHandler := handlers.NewCustomerHandler(customerUsecase)

	return &handlers.Handlers{
		CustomerHandler: customerHandler,
	}
}

func startServer(engine *gin.Engine) {
	err := engine.Run(os.Getenv(config.EnvServerURL))
	if err != nil {
		panic("unable to start http server")
	}
}

package routes

import (
	"genesis_test_case/src/pkg/delivery/api"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/delivery/http/presenters"
	"genesis_test_case/src/pkg/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App) error {
	repos, err := api.CreateRepositories()
	if err != nil {
		return err
	}
	ucases, err := api.CreateUsecases(repos)
	if err != nil {
		return err
	}
	handlers, err := createHandlers(ucases)
	if err != nil {
		return err
	}

	middleware.FiberMiddleware(app)
	InitPublicRoutes(app, handlers)

	return nil
}

func createHandlers(usecases *usecases.Usecases) (*Handlers, error) {
	cryptoMailingUsecases := &http.CryptoMailingUsecases{
		Mailing:      usecases.CryptoMailing,
		Subscription: usecases.Subscription,
	}
	presenter := presenters.NewPresenterJSON()
	mailingHandler := http.NewMailingHandler(cryptoMailingUsecases, presenter)
	rateHandler := http.NewConfigRateHandler(usecases.CryptoExchanger, presenter)

	return &Handlers{
		Mailing: mailingHandler,
		Rate:    rateHandler,
	}, nil
}

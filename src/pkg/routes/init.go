package routes

import (
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/repository"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/platform/gmail_api"

	"github.com/gofiber/fiber/v2"
)

func initHandler() (*http.MailingHandler, error) {
	gmailService, err := gmail_api.GetGmailService()
	if err != nil {
		return nil, err
	}
	repos := repository.NewRepositories(gmailService)
	usecases := usecase.NewUsecases(repos)
	handler := http.NewMailingHandler(usecases)

	return handler, nil
}

func InitRoutes(app *fiber.App) error {
	handler, err := initHandler()
	if err != nil {
		return err
	}

	middleware.FiberMiddleware(app)
	InitPublicRoutes(app, handler)

	return nil
}

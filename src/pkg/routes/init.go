package routes

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/persistence/crypto/banners"
	"genesis_test_case/src/pkg/persistence/crypto/exchangers"
	"genesis_test_case/src/pkg/persistence/mailing"
	storage "genesis_test_case/src/pkg/persistence/storage/csv"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/platform/gmail_api"
	"os"

	"github.com/gofiber/fiber/v2"
)

func createRepositories() (*usecase.Repositories, error) {
	gmailService, err := gmail_api.GetGmailService()
	if err != nil {
		return nil, err
	}
	csvStorage := storage.NewCsvEmaiStorage(os.Getenv(config.EnvStorageFilePath))
	mailingGmailRepository := mailing.NewGmailRepository(gmailService)
	cryptoCoinbaseRepository := exchangers.NewCryptoCoinbaseRepository(
		os.Getenv(config.EnvCoinbaseApiExchangeUrl),
		os.Getenv(config.EnvCoinbaseApiCandlesUrl),
	)
	cryptoBannerBearRepository := banners.NewCryptoBannerBearRepository(
		os.Getenv(config.EnvBannerApiToken),
		os.Getenv(config.EnvBannerApiUrl),
		os.Getenv(config.EnvCryptoBannerTemplate),
	)

	return &usecase.Repositories{
		Banner:    cryptoBannerBearRepository,
		Exchanger: cryptoCoinbaseRepository,
		Storage:   csvStorage,
		Mailer:    mailingGmailRepository,
	}, nil
}

func createUsecases(repos *usecase.Repositories) *http.Usecases {
	cryptoMailingRepositories := &usecase.CryptoMailingRepositories{
		Repositories: *repos,
	}
	BTCUAHPair := &domain.CurrencyPair{
		BaseCurrency:  os.Getenv(config.EnvBaseCurrency),
		QuoteCurrency: os.Getenv(config.EnvQuoteCurrency),
	}
	cryptoMailignUsecase := usecase.NewCryptoMailingUsecase(
		os.Getenv(config.EnvCryptoHtmlMessagePath),
		BTCUAHPair,
		cryptoMailingRepositories,
	)
	cryptoExchangeUsecase := usecase.NewCryptoExchangeUsecase(
		BTCUAHPair,
		repos.Exchanger,
	)
	subscriptionUsecase := usecase.NewSubscriptionUsecase(
		repos.Storage,
	)

	return &http.Usecases{
		Subscription:    subscriptionUsecase,
		CryptoMailing:   cryptoMailignUsecase,
		CryptoExchanger: cryptoExchangeUsecase,
	}
}

func initHandler() (*http.MailingHandler, error) {
	repos, err := createRepositories()
	if err != nil {
		return nil, err
	}
	usecases := createUsecases(repos)
	cryptoMailingUsecases := &http.CryptoMailingUsecases{
		Exchange:     usecases.CryptoExchanger,
		Mailing:      usecases.CryptoMailing,
		Subscription: usecases.Subscription,
	}
	handler := http.NewMailingHandler(cryptoMailingUsecases)

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

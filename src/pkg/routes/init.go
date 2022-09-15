package routes

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/persistence/crypto/banners"
	"genesis_test_case/src/pkg/persistence/crypto/cache"
	"genesis_test_case/src/pkg/persistence/crypto/exchangers"
	"genesis_test_case/src/pkg/persistence/mailing"
	storage "genesis_test_case/src/pkg/persistence/storage/csv"
	"genesis_test_case/src/pkg/persistence/storage/redis"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/platform/gmail_api"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func createRepositories() (*usecase.Repositories, error) {
	gmailService, err := gmail_api.GetGmailService()
	if err != nil {
		return nil, err
	}
	csvStorage := storage.NewCsvEmaiStorage(os.Getenv(config.EnvStorageFilePath))
	mailingGmailRepository := mailing.NewGmailRepository(gmailService)
	cryptoBannerBearRepository := banners.NewCryptoBannerBearRepository(
		os.Getenv(config.EnvBannerApiToken),
		os.Getenv(config.EnvBannerApiUrl),
		os.Getenv(config.EnvCryptoBannerTemplate),
	)

	return &usecase.Repositories{
		Banner:  cryptoBannerBearRepository,
		Storage: csvStorage,
		Mailer:  mailingGmailRepository,
	}, nil
}

func setupUsecases(repos *usecase.Repositories) (*http.Usecases, error) {
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

	cryptoCache, err := setupCryptoCache()
	if err != nil {
		return nil, err
	}

	configuredExchanger := getConfiguredExchanger()

	cryptoExchangeUsecase := usecase.NewCryptoExchangeUsecase(
		BTCUAHPair,
		configuredExchanger,
		cryptoCache,
	)

	subscriptionUsecase := usecase.NewSubscriptionUsecase(
		repos.Storage,
	)

	return &http.Usecases{
		Subscription:    subscriptionUsecase,
		CryptoMailing:   cryptoMailignUsecase,
		CryptoExchanger: cryptoExchangeUsecase,
	}, nil
}

func setupCryptoCache() (usecase.CryptoCache, error) {
	cryptoCacheDB, err := strconv.Atoi(os.Getenv(config.CryptoCacheDB))
	if err != nil {
		return nil, err
	}

	cacheExpiresMins, err := strconv.Atoi(os.Getenv(config.CryptoCacheExpiresMins))
	if err != nil {
		return nil, err
	}

	cacheProvider := redis.NewRedisCache(
		os.Getenv(config.CryptoCacheHost),
		cryptoCacheDB,
		time.Duration(cacheExpiresMins)*time.Minute,
	)

	return cache.NewCryptoCache(cacheProvider), nil
}

func getConfiguredExchanger() usecase.ExchangeProvider {
	coinapiExchangerNode := exchangers.CoinApiProviderFactory{}.CreateExchangeProviderNode()
	coinbaseExchangerNode := exchangers.CoinbaseProviderFactory{}.CreateExchangeProviderNode()
	nomicsExchangerNode := exchangers.NomicsProviderFactory{}.CreateRateService()

	chain := usecase.NewExchangersChain()
	if chain.RegisterExchanger(
		config.CoinAPIExchangerName,
		coinapiExchangerNode,
		coinbaseExchangerNode,
	)
	chain.RegisterExchanger(
		config.CoinbaseExchangerName,
		coinbaseExchangerNode,
		nomicsExchangerNode,
	)
	chain.RegisterExchanger(
		config.NomicsExchangerName,
		nomicsExchangerNode,
		nil,
	)

	return chain.GetExchanger(
		os.Getenv(config.EnvDefaultExchangerName),
	)
}

func initHandler() (*http.MailingHandler, error) {
	repos, err := createRepositories()
	if err != nil {
		return nil, err
	}
	usecases, err := setupUsecases(repos)
	if err != nil {
		return nil, err
	}
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

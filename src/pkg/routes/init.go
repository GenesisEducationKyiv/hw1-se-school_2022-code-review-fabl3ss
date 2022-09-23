package routes

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/loggers"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/persistence/crypto"
	"genesis_test_case/src/pkg/persistence/crypto/banners"
	"genesis_test_case/src/pkg/persistence/crypto/charts"
	"genesis_test_case/src/pkg/persistence/crypto/exchangers"
	"genesis_test_case/src/pkg/persistence/mailing"
	storage "genesis_test_case/src/pkg/persistence/storage/csv"
	"genesis_test_case/src/pkg/persistence/storage/redis"
	"genesis_test_case/src/pkg/usecase"
	exchangeUsecase "genesis_test_case/src/pkg/usecase/exchange"
	mailingUsecase "genesis_test_case/src/pkg/usecase/mailing"
	subscriptionUsecase "genesis_test_case/src/pkg/usecase/subscription"
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
	cryptobannerBearProvidersitory := banners.BannerBearProviderFactory{}.CreateBannerProvider()
	exchangeProvider := exchangers.CoinApiProviderFactory{}.CreateExchangeProvider()
	chartProvider := charts.CoinbaseProviderFactory{}.CreateChartProvider()
	return &usecase.Repositories{
		Banner:    cryptobannerBearProvidersitory,
		Storage:   csvStorage,
		Mailer:    mailingGmailRepository,
		Exchanger: exchangeProvider,
		Chart:     chartProvider,
	}, nil
}

func createUsecases(repos *usecase.Repositories) (*http.Usecases, error) {
	cryptoMailingRepositories := &usecase.CryptoMailingRepositories{
		Repositories: *repos,
	}
	BTCUAHPair := domain.NewCurrencyPair(
		os.Getenv(config.EnvBaseCurrency),
		os.Getenv(config.EnvQuoteCurrency),
	)
	cryptoMailingBecause := mailingUsecase.NewCryptoMailingUsecase(
		os.Getenv(config.EnvCryptoHtmlMessagePath),
		BTCUAHPair,
		cryptoMailingRepositories,
	)

	cryptoCache, err := setupCryptoCache()
	if err != nil {
		return nil, err
	}

	configuredExchanger := getConfiguredExchanger()

	cryptoExchangeUsecase := exchangeUsecase.NewCryptoExchangeUsecase(
		configuredExchanger,
		cryptoCache,
	)

	subscriptionUsecase := subscriptionUsecase.NewSubscriptionUsecase(
		repos.Storage,
	)

	return &http.Usecases{
		Subscription:    subscriptionUsecase,
		CryptoMailing:   cryptoMailingBecause,
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

	return crypto.NewCryptoCache(cacheProvider), nil
}

func getConfiguredExchanger() usecase.ExchangeProvider {
	logger := loggers.NewZapLogger(os.Getenv(config.EnvLogPath))
	cryptoLogger := crypto.NewCryptoLogger(logger)

	coinapiExchanger := exchangers.CoinApiProviderFactory{}.CreateExchangeProvider()
	coinbaseExchanger := exchangers.CoinbaseProviderFactory{}.CreateExchangeProvider()
	nomicsExchanger := exchangers.NomicsProviderFactory{}.CreateExchangeProvider()

	loggingCoinapiExchanger := exchangers.NewLoggingExchanger(coinapiExchanger, cryptoLogger)
	loggingCoinbaseExchanger := exchangers.NewLoggingExchanger(coinbaseExchanger, cryptoLogger)
	loggingNomicsExchanger := exchangers.NewLoggingExchanger(nomicsExchanger, cryptoLogger)

	coinapiExchangerNode := exchangers.NewExchangerNode(loggingCoinapiExchanger)
	coinbaseExchangerNode := exchangers.NewExchangerNode(loggingCoinbaseExchanger)
	nomicsExchangerNode := exchangers.NewExchangerNode(loggingNomicsExchanger)

	chain := exchangeUsecase.NewExchangersChain()
	chain.RegisterExchanger(
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

func createHandlers(usecases *http.Usecases) (*Handlers, error) {
	cryptoMailingUsecases := &http.CryptoMailingUsecases{
		Exchange:     usecases.CryptoExchanger,
		Mailing:      usecases.CryptoMailing,
		Subscription: usecases.Subscription,
	}
	mailingHandler := http.NewMailingHandler(cryptoMailingUsecases)
	rateHandler := http.NewConfigRateHandler(usecases.CryptoExchanger)

	return &Handlers{
		Mailing: mailingHandler,
		Rate:    rateHandler,
	}, nil
}

func InitRoutes(app *fiber.App) error {
	repos, err := createRepositories()
	if err != nil {
		return err
	}
	usecases, err := createUsecases(repos)
	if err != nil {
		return err
	}
	handlers, err := createHandlers(usecases)
	if err != nil {
		return err
	}

	middleware.FiberMiddleware(app)
	InitPublicRoutes(app, handlers)

	return nil
}

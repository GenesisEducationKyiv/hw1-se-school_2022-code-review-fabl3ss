package api

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/loggers"
	"genesis_test_case/src/pkg/application"
	exchangeUsecase "genesis_test_case/src/pkg/application/exchange"
	mailingUsecase "genesis_test_case/src/pkg/application/mailing"
	subscriptionUsecase "genesis_test_case/src/pkg/application/subscription"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecases"
	"genesis_test_case/src/pkg/persistence/crypto"
	"genesis_test_case/src/pkg/persistence/crypto/banners"
	"genesis_test_case/src/pkg/persistence/crypto/charts"
	"genesis_test_case/src/pkg/persistence/crypto/exchangers"
	"genesis_test_case/src/pkg/persistence/mailing"
	storage "genesis_test_case/src/pkg/persistence/storage/csv"
	"genesis_test_case/src/pkg/persistence/storage/redis"
	"genesis_test_case/src/platform/gmail_api"
	"os"
	"strconv"
	"time"
)

func CreateUsecases(repos *application.Repositories) (*usecases.Usecases, error) {
	cryptoMailingRepositories := &application.CryptoMailingRepositories{
		Repositories: *repos,
	}
	BTCUAHPair := models.NewCurrencyPair(
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

	return &usecases.Usecases{
		Subscription:    subscriptionUsecase,
		CryptoMailing:   cryptoMailingBecause,
		CryptoExchanger: cryptoExchangeUsecase,
	}, nil
}

func CreateRepositories() (*application.Repositories, error) {
	gmailService, err := gmail_api.GetGmailService()
	if err != nil {
		return nil, err
	}
	csvStorage := storage.NewCsvEmaiStorage(os.Getenv(config.EnvStorageFilePath))
	mailingGmailRepository := mailing.NewGmailRepository(gmailService)
	cryptobannerBearProvidersitory := banners.BannerBearProviderFactory{}.CreateBannerProvider()
	exchangeProvider := exchangers.CoinApiProviderFactory{}.CreateExchangeProvider()
	chartProvider := charts.CoinbaseProviderFactory{}.CreateChartProvider()
	return &application.Repositories{
		Banner:    cryptobannerBearProvidersitory,
		Storage:   csvStorage,
		Mailer:    mailingGmailRepository,
		Exchanger: exchangeProvider,
		Chart:     chartProvider,
	}, nil
}
func setupCryptoCache() (application.CryptoCache, error) {
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

func getConfiguredExchanger() application.ExchangeProvider {
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
	if err := chain.RegisterExchanger(
		config.CoinAPIExchangerName,
		coinapiExchangerNode,
		coinbaseExchangerNode,
	); err != nil {
		return nil
	}
	if err := chain.RegisterExchanger(
		config.CoinbaseExchangerName,
		coinbaseExchangerNode,
		nomicsExchangerNode,
	); err != nil {
		return nil
	}
	if err := chain.RegisterExchanger(
		config.NomicsExchangerName,
		nomicsExchangerNode,
		nil,
	); err != nil {
		return nil
	}

	return chain.GetExchanger(
		os.Getenv(config.EnvDefaultExchangerName),
	)
}

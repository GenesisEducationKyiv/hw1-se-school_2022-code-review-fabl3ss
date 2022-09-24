package usecases

type CryptoMailingUsecase interface {
	SendCurrencyRate() ([]string, error)
}

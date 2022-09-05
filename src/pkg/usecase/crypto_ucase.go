package usecase

import (
	"strconv"

	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/repository"
)

type cryptoUsecase struct {
	repos *repository.Repositories
}

func NewCryptoUsecase(r *repository.Repositories) domain.CryptoUsecase {
	return &cryptoUsecase{
		repos: r,
	}
}

func (c *cryptoUsecase) GetConfigCurrencyRate(cfg *config.Config) (int, error) {
	rate, err := c.repos.Crypto.GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(rate.Price)
}

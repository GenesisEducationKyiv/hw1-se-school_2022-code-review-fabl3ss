package exchangers

import (
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
)

type exchangerNode struct {
	exchanger usecase.ExchangeProvider
	next      usecase.ExchangeProvider
}

func NewExchangerNode(exc usecase.ExchangeProvider) usecase.ExchangeProviderNode {
	return &exchangerNode{
		exchanger: exc,
	}
}

func (c *exchangerNode) GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	rate, err := c.exchanger.GetCurrencyRate(pair)
	if err != nil && c.next != nil {
		return c.next.GetCurrencyRate(pair)
	}

	return rate, nil
}

func (c *exchangerNode) SetNext(service usecase.ExchangeProviderNode) {
	c.next = service
}

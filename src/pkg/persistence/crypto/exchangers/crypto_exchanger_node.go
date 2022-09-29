package exchangers

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
)

type exchangerNode struct {
	exchanger application.ExchangeProvider
	next      application.ExchangeProvider
}

func NewExchangerNode(exc application.ExchangeProvider) application.ExchangeProviderNode {
	return &exchangerNode{
		exchanger: exc,
	}
}

func (c *exchangerNode) GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error) {
	rate, err := c.exchanger.GetCurrencyRate(pair)
	if err != nil && c.next != nil {
		return c.next.GetCurrencyRate(pair)
	}

	return rate, nil
}

func (c *exchangerNode) SetNext(service application.ExchangeProviderNode) {
	c.next = service
}

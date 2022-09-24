package usecase

import (
	"genesis_test_case/src/pkg/application"
	myerr "genesis_test_case/src/pkg/types/errors"
)

type exchangersChain struct {
	exchangers map[string]application.ExchangeProviderNode
}

func NewExchangersChain() application.ExchangersChain {
	return &exchangersChain{
		exchangers: make(map[string]application.ExchangeProviderNode),
	}
}

func (e *exchangersChain) RegisterExchanger(name string, exchanger, next application.ExchangeProviderNode) error {
	if len(name) < 1 || exchanger == nil {
		return myerr.ErrInvalidInput
	}

	e.exchangers[name] = exchanger
	e.exchangers[name].SetNext(next)

	return nil
}

func (e *exchangersChain) GetExchanger(name string) application.ExchangeProvider {
	return e.exchangers[name]
}

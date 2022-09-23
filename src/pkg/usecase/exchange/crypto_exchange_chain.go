package usecase

import (
	myerr "genesis_test_case/src/pkg/types/errors"
	"genesis_test_case/src/pkg/usecase"
)

type exchangersChain struct {
	exchangers map[string]usecase.ExchangeProviderNode
}

func NewExchangersChain() usecase.ExchangersChain {
	return &exchangersChain{
		exchangers: make(map[string]usecase.ExchangeProviderNode),
	}
}

func (e *exchangersChain) RegisterExchanger(name string, exchanger, next usecase.ExchangeProviderNode) error {
	if len(name) < 1 || exchanger == nil {
		return myerr.ErrInvalidInput
	}

	e.exchangers[name] = exchanger
	e.exchangers[name].SetNext(next)

	return nil
}

func (e *exchangersChain) GetExchanger(name string) usecase.ExchangeProvider {
	return e.exchangers[name]
}

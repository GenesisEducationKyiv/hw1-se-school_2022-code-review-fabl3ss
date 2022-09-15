package usecase

import (
	myerr "genesis_test_case/src/pkg/types/errors"
)

type exchangersChain struct {
	exchangers map[string]ExchangeProviderNode
}

func NewExchangersChain() ExchangersChain {
	return &exchangersChain{
		exchangers: make(map[string]ExchangeProviderNode),
	}
}

func (e *exchangersChain) RegisterExchanger(name string, exchanger, next ExchangeProviderNode) error {
	if len(name) < 1 || exchanger == nil {
		return myerr.ErrInvalidInput
	}

	e.exchangers[name] = exchanger
	e.exchangers[name].SetNext(next)

	return nil
}

func (e *exchangersChain) GetExchanger(name string) ExchangeProvider {
	return e.exchangers[name]
}

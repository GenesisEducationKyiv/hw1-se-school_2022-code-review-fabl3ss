package models

type CurrencyPair struct {
	baseCurrency  string
	quoteCurrency string
}

func NewCurrencyPair(base string, quote string) *CurrencyPair {
	return &CurrencyPair{
		baseCurrency:  base,
		quoteCurrency: quote,
	}
}

func (c *CurrencyPair) GetBaseCurrency() string {
	return c.baseCurrency
}

func (c *CurrencyPair) GetQuoteCurrency() string {
	return c.quoteCurrency
}

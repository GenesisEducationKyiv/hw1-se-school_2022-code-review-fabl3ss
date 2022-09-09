package domain

type CurrencyPair struct {
	BaseCurrency  string
	QuoteCurrency string
}

type CurrencyRate struct {
	CurrencyPair
	Price float64
}

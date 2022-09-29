package config

const (
	NomicsExchangerTemplateURL   string = "https://api.nomics.com/v1/currencies/ticker?key=%v&ids=%v&interval=1d&convert=%v"
	CoinAPIExchangerTemplateURL  string = "https://rest.coinapi.io/v1/exchangerate/%v/%v"
	CoinbaseExchangerTemplateURL string = "https://api.coinbase.com/v2/prices/%s-%s/spot"
)

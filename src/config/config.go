package config

import (
	"os"
	"sync"
)

type Config struct {
	ServerURL           string
	ServerReadTimeout   string
	BaseCurrency        string
	QuoteCurrency       string
	CryptoApiFormatUrl  string
	CryptoApiCandlesUrl string
	StorageFile         string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		cfg = Config{
			ServerURL:           os.Getenv(ServerUrl),
			ServerReadTimeout:   os.Getenv(ServerReadTimeout),
			CryptoApiFormatUrl:  os.Getenv(CryptoApiFormatUrl),
			CryptoApiCandlesUrl: os.Getenv(CryptoApiCandlesUrl),
			BaseCurrency:        os.Getenv(BaseCurrency),
			QuoteCurrency:       os.Getenv(QuoteCurrency),
			StorageFile:         os.Getenv(StorageFilePath),
		}
	})
	return &cfg
}

package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestConfigGet(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	var (
		ServerUrl             string = os.Getenv(EnvServerUrl)
		ServerReadTimeout     string = os.Getenv(EnvServerReadTimeout)
		BaseCurrency          string = os.Getenv(EnvBaseCurrency)
		QuoteCurrency         string = os.Getenv(EnvQuoteCurrency)
		StorageFilePath       string = os.Getenv(EnvStorageFilePath)
		GmailTokenPath        string = os.Getenv(EnvGmailTokenPath)
		GmailCredentialsPath  string = os.Getenv(EnvGmailCredentialsPath)
		BannerApiToken        string = os.Getenv(EnvBannerApiToken)
		BannerApiUrl          string = os.Getenv(EnvBannerApiUrl)
		CryptoBannerTemplate  string = os.Getenv(EnvCryptoBannerTemplate)
		CryptoHtmlMessagePath string = os.Getenv(EnvCryptoHtmlMessagePath)
		CryptoApiExchangeUrl  string = os.Getenv(EnvCoinbaseApiExchangeUrl)
		CryptoApiCandlesUrl   string = os.Getenv(EnvCoinbaseApiCandlesUrl)
	)

	require.NotEmpty(t, ServerUrl)
	require.NotEmpty(t, ServerReadTimeout)
	require.NotEmpty(t, BaseCurrency)
	require.NotEmpty(t, QuoteCurrency)
	require.NotEmpty(t, StorageFilePath)
	require.NotEmpty(t, GmailTokenPath)
	require.NotEmpty(t, GmailCredentialsPath)
	require.NotEmpty(t, BannerApiToken)
	require.NotEmpty(t, BannerApiUrl)
	require.NotEmpty(t, CryptoBannerTemplate)
	require.NotEmpty(t, CryptoHtmlMessagePath)
	require.NotEmpty(t, CryptoApiExchangeUrl)
	require.NotEmpty(t, CryptoApiCandlesUrl)
}

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
		ServerUrl             = os.Getenv(EnvServerUrl)
		ServerReadTimeout     = os.Getenv(EnvServerReadTimeout)
		BaseCurrency          = os.Getenv(EnvBaseCurrency)
		QuoteCurrency         = os.Getenv(EnvQuoteCurrency)
		StorageFilePath       = os.Getenv(EnvCsvStoragePath)
		GmailTokenPath        = os.Getenv(EnvGmailTokenPath)
		GmailCredentialsPath  = os.Getenv(EnvGmailCredentialsPath)
		BannerApiToken        = os.Getenv(EnvBannerApiToken)
		CryptoBannerTemplate  = os.Getenv(EnvCryptoBannerTemplate)
		CryptoHtmlMessagePath = os.Getenv(EnvCryptoHtmlMessagePath)
		DefaultExchangerName  = os.Getenv(EnvDefaultExchangerName)
	)

	require.NotEmpty(t, ServerUrl)
	require.NotEmpty(t, ServerReadTimeout)
	require.NotEmpty(t, BaseCurrency)
	require.NotEmpty(t, QuoteCurrency)
	require.NotEmpty(t, StorageFilePath)
	require.NotEmpty(t, GmailTokenPath)
	require.NotEmpty(t, GmailCredentialsPath)
	require.NotEmpty(t, BannerApiToken)
	require.NotEmpty(t, CryptoBannerTemplate)
	require.NotEmpty(t, CryptoHtmlMessagePath)
	require.NotEmpty(t, DefaultExchangerName)
}

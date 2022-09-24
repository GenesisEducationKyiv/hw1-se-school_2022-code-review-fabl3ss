package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchCryptoExchangers(t *testing.T) {
	archtest.Package(t, cryptoExchangersPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoChartsPackage,
		cryptoBannersPackage,
		mailingPackage,
		storageCsvPackage,
		storageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchCryptoBanners(t *testing.T) {
	archtest.Package(t, cryptoBannersPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoChartsPackage,
		cryptoExchangersPackage,
		mailingPackage,
		storageCsvPackage,
		storageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchCryptoCharts(t *testing.T) {
	archtest.Package(t, cryptoBannersPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoBannersPackage,
		cryptoExchangersPackage,
		mailingPackage,
		storageCsvPackage,
		storageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchMailing(t *testing.T) {
	archtest.Package(t, mailingPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoBannersPackage,
		cryptoExchangersPackage,
		storageCsvPackage,
		storageRedisPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchStorageCsv(t *testing.T) {
	archtest.Package(t, storageCsvPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoBannersPackage,
		cryptoExchangersPackage,
		mailingPackage,
		storageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchStorageRedis(t *testing.T) {
	archtest.Package(t, storageCsvPackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoBannersPackage,
		cryptoExchangersPackage,
		mailingPackage,
		storageCsvPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestCryptoExchangersHaveTests(t *testing.T) {
	archtest.Package(t, cryptoExchangersPackage).IncludeTests()
}

package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchCryptoExchangers(t *testing.T) {
	archtest.Package(t, persistenceCryptoExchangersPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoBannersPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchCryptoBanners(t *testing.T) {
	archtest.Package(t, persistenceCryptoBannersPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoExchangersPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchCryptoCharts(t *testing.T) {
	archtest.Package(t, persistenceCryptoBannersPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchMailing(t *testing.T) {
	archtest.Package(t, persistenceMailingPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchStorageCsv(t *testing.T) {
	archtest.Package(t, persistenceStorageCsvPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		persistenceMailingPackage,
		persistenceStorageRedisPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestArchStorageRedis(t *testing.T) {
	archtest.Package(t, persistenceStorageCsvPackage).ShouldNotDependOn(
		persistenceCryptoPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		gmailPlatformPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		httpHandlersPackage,
	)
}

func TestCryptoExchangersHaveTests(t *testing.T) {
	archtest.Package(t, persistenceCryptoExchangersPackage).IncludeTests()
}

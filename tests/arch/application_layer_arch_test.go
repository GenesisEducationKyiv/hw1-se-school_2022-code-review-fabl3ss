package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchApplication(t *testing.T) {
	archtest.Package(t, applicationExchangePackage).ShouldNotDependOn(
		applicationMailingPackage,
		applicationExchangePackage,
		applicationSubscriptionPackage,
		persistenceCryptoPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		httpRoutesPackage,
		httpHandlersPackage,
		gmailPlatformPackage,
	)
}

func TestArchApplicationExchange(t *testing.T) {
	archtest.Package(t, applicationExchangePackage).ShouldNotDependOn(
		applicationMailingPackage,
		applicationSubscriptionPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		persistenceCryptoPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		httpRoutesPackage,
		httpHandlersPackage,
		gmailPlatformPackage,
	)
}

func TestArchApplicationMailing(t *testing.T) {
	archtest.Package(t, applicationExchangePackage).ShouldNotDependOn(
		applicationExchangePackage,
		applicationSubscriptionPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		persistenceCryptoPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		httpRoutesPackage,
		httpHandlersPackage,
		gmailPlatformPackage,
	)
}

func TestArchApplicationSubscription(t *testing.T) {
	archtest.Package(t, applicationSubscriptionPackage).ShouldNotDependOn(
		applicationExchangePackage,
		applicationMailingPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		persistenceCryptoPackage,
		persistenceCryptoExchangersPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoBannersPackage,
		httpRoutesPackage,
		httpHandlersPackage,
		gmailPlatformPackage,
	)
}

func TestApplicationExchangeHaveTests(t *testing.T) {
	archtest.Package(t, applicationExchangePackage).IncludeTests()
}

func TestApplicationMailingHaveTests(t *testing.T) {
	archtest.Package(t, applicationMailingPackage).IncludeTests()
}

func TestApplicationSubscriptionHaveTests(t *testing.T) {
	archtest.Package(t, applicationSubscriptionPackage).IncludeTests()
}

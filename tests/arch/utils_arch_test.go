package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestUtilsHaveNoDependencies(t *testing.T) {
	archtest.Package(t, utilsPackage).ShouldNotDependOn(
		modelsPackage,
		applicationPackage,
		applicationExchangePackage,
		applicationMailingPackage,
		applicationSubscriptionPackage,
		httpHandlersPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoExchangersPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoPackage,
		persistenceMailingPackage,
		persistenceStorageCsvPackage,
		persistenceStorageRedisPackage,
		loggersPackage,
		configPackage,
		gmailPlatformPackage,
	)
}

func TestUtilsHaveTests(t *testing.T) {
	archtest.Package(t, utilsPackage).IncludeTests()
}

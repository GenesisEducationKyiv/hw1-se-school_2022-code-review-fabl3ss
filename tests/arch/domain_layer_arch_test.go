package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestModelsHaveNoDependencies(t *testing.T) {
	archtest.Package(t, modelsPackage).ShouldNotDependOn(
		httpRoutesPackage,
		httpHandlersPackage,
		applicationMailingPackage,
		applicationExchangePackage,
		applicationSubscriptionPackage,
		persistenceCryptoBannersPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoExchangersPackage,
		persistenceCryptoChartsPackage,
		persistenceCryptoPackage,
		loggersPackage,
		configPackage,
		gmailPlatformPackage,
	)
}

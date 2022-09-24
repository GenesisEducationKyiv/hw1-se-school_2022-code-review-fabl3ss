package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestUtilsHaveNoDependencies(t *testing.T) {
	archtest.Package(t, utilsPackage).ShouldNotDependOn(
		modelsPackage,
		usecasePackage,
		httpHandlersPackage,
		httpRoutesPackage,
		httpPresentersPackage,
		cryptoBannersPackage,
		cryptoExchangersPackage,
		cryptoChartsPackage,
		cryptoPackage,
		mailingPackage,
		storageCsvPackage,
		storageRedisPackage,
		loggersPackage,
		configPackage,
		gmailPlatformPackage,
	)
}

func TestUtilsHaveTests(t *testing.T) {
	archtest.Package(t, utilsPackage).IncludeTests()
}

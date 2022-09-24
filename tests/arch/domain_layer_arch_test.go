package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchDomainModels(t *testing.T) {
	archtest.Package(t, modelsPackage).ShouldNotDependOn(
		usecasePackage,
		httpRoutesPackage,
		httpHandlersPackage,
		cryptoBannersPackage,
		cryptoChartsPackage,
		cryptoExchangersPackage,
		cryptoChartsPackage,
		cryptoPackage,
		loggersPackage,
		configPackage,
		gmailPlatformPackage,
	)
}

func TestArchDomainUsecases(t *testing.T) {
	archtest.Package(t, usecasePackage).ShouldNotDependOn(
		cryptoPackage,
		cryptoExchangersPackage,
		cryptoChartsPackage,
		cryptoBannersPackage,
		mailingPackage,
		storageCsvPackage,
		storageRedisPackage,
		httpRoutesPackage,
		httpHandlersPackage,
		gmailPlatformPackage,
	)
}

func TestDomainUsecasesHaveTests(t *testing.T) {
	archtest.Package(t, usecasePackage).IncludeTests()
}

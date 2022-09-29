package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestDomainLayerHaveNoDependencies(t *testing.T) {
	archtest.Package(t, domainLayer).ShouldNotDependOn(
		utilsPackage,
		configPackage,
		platformLayer,
		loggersPackage,
		applicationLayer,
		persistenceLayer,
		presentationLayer,
	)
}

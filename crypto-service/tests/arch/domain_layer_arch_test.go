package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
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

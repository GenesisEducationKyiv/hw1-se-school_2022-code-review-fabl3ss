package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestPlatformLayerHaveNoDependencies(t *testing.T) {
	archtest.Package(t, domainLayer).ShouldNotDependOn(
		utilsPackage,
		configPackage,
		domainLayer,
		loggersPackage,
		applicationLayer,
		persistenceLayer,
		presentationLayer,
	)
}

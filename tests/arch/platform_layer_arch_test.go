package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
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

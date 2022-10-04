package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestUtilsHaveNoDependencies(t *testing.T) {
	archtest.Package(t, utilsPackage).ShouldNotDependOn(
		domainLayer,
		platformLayer,
		configPackage,
		loggersPackage,
		applicationLayer,
		persistenceLayer,
		presentationLayer,
	)
}

func TestUtilsHaveTests(t *testing.T) {
	archtest.Package(t, utilsPackage).IncludeTests()
}

package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
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

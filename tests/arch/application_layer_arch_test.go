package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchApplicationLayer(t *testing.T) {
	archtest.Package(t, applicationLayer).ShouldNotDependOn(
		platformLayer,
		persistenceLayer,
		presentationLayer,
	)
}

func TestApplicationLayerHaveTests(t *testing.T) {
	archtest.Package(t, applicationLayer).IncludeTests()
}

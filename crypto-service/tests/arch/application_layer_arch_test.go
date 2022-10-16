package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
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

package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestArchPersistenceLayer(t *testing.T) {
	archtest.Package(t, persistenceLayer).ShouldNotDependOn(
		presentationLayer,
	)
}

func TestPersistenceLayerHaveTests(t *testing.T) {
	archtest.Package(t, persistenceLayer).IncludeTests()
}

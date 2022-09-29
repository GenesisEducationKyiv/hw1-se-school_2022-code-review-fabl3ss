package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchPersistenceLayer(t *testing.T) {
	archtest.Package(t, persistenceLayer).ShouldNotDependOn(
		presentationLayer,
	)
}

func TestPersistenceLayerHaveTests(t *testing.T) {
	archtest.Package(t, persistenceLayer).IncludeTests()
}

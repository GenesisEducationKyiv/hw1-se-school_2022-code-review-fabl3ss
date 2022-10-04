package arch

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestArchPresentationLayer(t *testing.T) {
	archtest.Package(t, presentationLayer).ShouldNotDependDirectlyOn(
		platformLayer,
		applicationLayer,
		persistenceLayer,
	)
}

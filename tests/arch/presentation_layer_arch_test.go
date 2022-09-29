package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchPresentationLayer(t *testing.T) {
	archtest.Package(t, presentationLayer).ShouldNotDependDirectlyOn(
		platformLayer,
		applicationLayer,
		persistenceLayer,
	)
}

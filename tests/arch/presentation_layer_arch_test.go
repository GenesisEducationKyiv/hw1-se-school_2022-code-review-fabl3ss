package arch

import (
	"github.com/matthewmcnew/archtest"
	"testing"
)

func TestArchHttpHandlers(t *testing.T) {
	archtest.Package(t, httpHandlersPackage).ShouldNotDependOn(
		httpPresentersPackage,
	)
}

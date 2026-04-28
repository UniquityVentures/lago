package lago

import (
	"github.com/UniquityVentures/lago/components"
	"github.com/UniquityVentures/lago/registry"
)

var RegistryPage *registry.Registry[components.PageInterface] = registry.NewRegistry[components.PageInterface]()

package lago

import (
	"github.com/UniquityVentures/lago/registry"
	"github.com/UniquityVentures/lago/views"
)

var RegistryLayer *registry.Registry[views.GlobalLayer] = registry.NewRegistry[views.GlobalLayer]()

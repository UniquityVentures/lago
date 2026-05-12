package lago

import (
	"github.com/UniquityVentures/lago/registry"
	"github.com/UniquityVentures/lago/views"
)

var RegistryLayer *registry.ImmutableRegistry[views.GlobalLayer] = &registry.ImmutableRegistry[views.GlobalLayer]{}

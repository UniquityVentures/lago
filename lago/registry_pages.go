package lago

import (
	"github.com/UniquityVentures/lago/components"
	"github.com/UniquityVentures/lago/registry"
)

var RegistryPage *registry.ImmutableRegistry[components.PageInterface] = &registry.ImmutableRegistry[components.PageInterface]{}

package lago

import (
	"github.com/UniquityVentures/lago/registry"
)

type Config interface {
	PostConfig()
}

var RegistryConfig *registry.ImmutableRegistry[Config] = &registry.ImmutableRegistry[Config]{}

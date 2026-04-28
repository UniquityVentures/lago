package lago

import (
	"github.com/UniquityVentures/lago/registry"
)

type Config interface {
	PostConfig()
}

var RegistryConfig *registry.Registry[Config] = registry.NewRegistry[Config]()

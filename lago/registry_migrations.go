package lago

import (
	"github.com/UniquityVentures/lago/registry"
)

var RegistryMigrations *registry.ImmutableRegistry[UsefulFilesystem] = &registry.ImmutableRegistry[UsefulFilesystem]{}

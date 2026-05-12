package lago

import (
	"github.com/UniquityVentures/lago/registry"
	"github.com/spf13/cobra"
)

type CommandFactory func(LagoConfig) *cobra.Command

var RegistryCommand *registry.ImmutableRegistry[CommandFactory] = &registry.ImmutableRegistry[CommandFactory]{}

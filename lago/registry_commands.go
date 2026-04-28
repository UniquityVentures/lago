package lago

import (
	"github.com/UniquityVentures/lago/registry"
	"github.com/spf13/cobra"
)

type CommandFactory func(LagoConfig) *cobra.Command

var RegistryCommand *registry.Registry[CommandFactory] = registry.NewRegistry[CommandFactory]()

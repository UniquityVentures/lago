package p_pwa

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

// GetPlugin returns registry contributions for [lago.BuildAllRegistries].
// Shell head registrations for the manifest link remain in init() (see views.go).
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse("/")
	if err != nil {
		log.Panic(err)
	}
	return registry.Pair[string, lago.Plugin]{
		Key: "p_pwa",
		Value: lago.Plugin{
			Type:        lago.PluginTypeAddon,
			Icon:        "cpu-chip",
			URL:         u,
			VerboseName: "PWA",
			Configs:     pluginConfigs(),
			Views:       pluginViews(),
			Routes:      pluginRoutes(),
		},
	}
}

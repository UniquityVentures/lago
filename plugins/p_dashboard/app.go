package p_dashboard

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

const AppUrl = "/dashboard/"

// GetPlugin returns the registry contributions for this plugin (views, pages, routes) for
// [lago.BuildAllRegistries]. Callers that assemble the full plugin list should include
// a pair with key "p_dashboard" and this value.
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lago.Plugin]{Key: "p_dashboard", Value: lago.Plugin{
		Type:        lago.PluginTypeApp,
		Icon:        "dashboard",
		URL:         u,
		VerboseName: "Dashboard",
		Views:       pluginViews(),
		Pages:       pluginPages(),
		Routes:      pluginRoutes(),
	},
	}
}

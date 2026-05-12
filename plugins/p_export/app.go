package p_export

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

const AppUrl = "/export/"

// GetPlugin returns the registry contributions for this plugin (views, pages, routes) for
// [lago.BuildAllRegistries]. Callers that assemble the full plugin list should include
// a pair with key "p_export" and this value.
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lago.Plugin]{Key: "p_export", Value: lago.Plugin{
		Type:        lago.PluginTypeApp,
		Icon:        "arrow-down-tray",
		URL:         u,
		VerboseName: "Export",
		Views:       pluginViews(),
		Pages:       pluginPages(),
		Routes:      pluginRoutes(),
	},
	}
}

package p_livereloading

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

// GetPlugin returns routes and metadata for [lago.BuildAllRegistries].
// Shell head snippet registration stays in init() (see pages.go).
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse("/")
	if err != nil {
		log.Panic(err)
	}
	return registry.Pair[string, lago.Plugin]{
		Key: "p_livereloading",
		Value: lago.Plugin{
			Type:        lago.PluginTypeAddon,
			Icon:        "arrow-path",
			URL:         u,
			VerboseName: "Live reload",
			Routes:      pluginRoutes(),
		},
	}
}

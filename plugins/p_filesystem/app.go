package p_filesystem

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

const AppUrl = "/filesystem/"

// GetPlugin returns the registry contributions for this plugin for [lago.BuildAllRegistries].
// Callers assembling the full plugin list should include a pair with key "p_filesystem" and this value.
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lago.Plugin]{Key: "p_filesystem", Value: lago.Plugin{
		Type:        lago.PluginTypeApp,
		Icon:        "folder",
		URL:         u,
		VerboseName: "Filesystem",
		Roles:       []string{"superuser", "admin"},
		Migrations:  pluginMigrations(),
		Views:       pluginViews(),
		Pages:       pluginPages(),
		Routes:      pluginRoutes(),
		Configs:     pluginConfigs(),
		Generators: pluginGenerators(),
	},
	}
}

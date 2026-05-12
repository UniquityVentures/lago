package p_users

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)
// GetPlugin returns the registry contributions for this plugin for [lago.BuildAllRegistries].
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse(AppUrl)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lago.Plugin]{
		Key: "p_users",
		Value: lago.Plugin{
			Type:             lago.PluginTypeApp,
			Icon:             "users",
			URL:              u,
			VerboseName:      "Users",
			Migrations:       pluginMigrations(),
			Views:            pluginViews(),
			Pages:            pluginPages(),
			Routes:           pluginRoutes(),
			Models:           pluginModels(),
			Generators:       pluginGenerators(),
			DBInitHooks:      pluginDBInitHooks(),
			Configs:          pluginAuthConfigs(),
			CommandFactories: pluginCommandFactories(),
		},
	}
}

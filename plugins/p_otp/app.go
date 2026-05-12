package p_otp

import (
	"log"
	"net/url"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

const AppURL = "/otp/preferences/"

// GetPlugin returns registry contributions for [lago.BuildAllRegistries].
func GetPlugin() registry.Pair[string, lago.Plugin] {
	u, err := url.Parse(AppURL)
	if err != nil {
		log.Panic(err)
	}

	return registry.Pair[string, lago.Plugin]{
		Key: "p_otp",
		Value: lago.Plugin{
			Type:        lago.PluginTypeApp,
			Icon:        "key",
			URL:         u,
			VerboseName: "OTP Preferences",
			Roles:       []string{""},
			Views:       pluginViews(),
			Pages:       pluginPages(),
			Routes:      pluginRoutes(),
			Models:      pluginModels(),
		},
	}
}

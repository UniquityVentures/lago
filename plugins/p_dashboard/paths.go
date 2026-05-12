package p_dashboard

import "github.com/UniquityVentures/lago/lago"

// pluginRoutes returns HTTP routes for this plugin. The dashboard is view-driven;
// add [lago.Route] entries here when this plugin exposes standalone handlers.
func pluginRoutes() lago.PluginFeatures[lago.Route] {
	return lago.PluginFeatures[lago.Route]{}
}

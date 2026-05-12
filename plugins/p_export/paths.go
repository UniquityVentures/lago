package p_export

import (
	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

func pluginRoutes() lago.PluginFeatures[lago.Route] {
	return lago.PluginFeatures[lago.Route]{
		Entries: []registry.Pair[string, lago.Route]{
			{Key: "export.PageRoute", Value: lago.Route{
				Path:    AppUrl,
				Handler: lago.NewDynamicView("export.PageView"),
			}},
			{Key: "export.DownloadRoute", Value: lago.Route{
				Path:    AppUrl + "download/",
				Handler: lago.NewDynamicView("export.DownloadView"),
			}},
		},
	}
}

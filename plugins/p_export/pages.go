package p_export

import (
	"github.com/UniquityVentures/lago/components"
	"github.com/UniquityVentures/lago/getters"
	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

func pluginPages() lago.PluginFeatures[components.PageInterface] {
	return lago.PluginFeatures[components.PageInterface]{
		Entries: []registry.Pair[string, components.PageInterface]{
			{Key: "export.Menu", Value: components.SidebarMenu{
				Title: getters.Static("Export"),
				Back: &components.SidebarMenuItem{
					Title: getters.Static("Back to All Apps"),
					Url:   lago.RoutePath("dashboard.AppsPage", nil),
				},
				Children: []components.PageInterface{
					components.SidebarMenuItem{
						Title:  getters.Static("XLSX Export"),
						Url:    lago.RoutePath("export.PageRoute", nil),
						Active: true,
					},
				},
			}},
			{Key: "export.Page", Value: &components.ShellScaffold{
				Sidebar: []components.PageInterface{
					lago.DynamicPage{Name: "export.Menu"},
				},
				Children: []components.PageInterface{
					exportPickerPage{},
				},
			}},
		},
	}
}

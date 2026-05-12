package p_users

import (
	"github.com/UniquityVentures/lago/components"
	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

func pluginPages() lago.PluginFeatures[components.PageInterface] {
	var entries []registry.Pair[string, components.PageInterface]
	entries = append(entries, pageEntriesMenus()...)
	entries = append(entries, pageEntriesFilters()...)
	entries = append(entries, pageEntriesTables()...)
	entries = append(entries, pageEntriesDetail()...)
	entries = append(entries, pageEntriesForms()...)
	entries = append(entries, pageEntriesAuth()...)
	entries = append(entries, pageEntriesSelection()...)
	entries = append(entries, pageEntriesRole()...)
	return lago.PluginFeatures[components.PageInterface]{Entries: entries}
}

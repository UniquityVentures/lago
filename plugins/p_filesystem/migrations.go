package p_filesystem

import (
	"embed"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

//go:embed migrations
var migrationsFS embed.FS

func pluginMigrations() lago.PluginFeatures[lago.UsefulFilesystem] {
	return lago.PluginFeatures[lago.UsefulFilesystem]{
		Entries: []registry.Pair[string, lago.UsefulFilesystem]{
			{Key: "p_filesystem.migrations", Value: migrationsFS},
		},
	}
}

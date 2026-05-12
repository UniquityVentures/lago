# Migrating plugins to `GetPlugin` and `BuildAllRegistries`

This guide describes how to move a plugin from **`init()`-time `Registry*.Register`** calls to a single **`GetPlugin()`** function whose contributions are merged by **`lago.BuildAllRegistries`**.

Referenced types live in **`lago/registry_plugins.go`** (`Plugin`, `PluginFeatures`) and **`lago/registry.go`** (`BuildAllRegistries`).

## Why migrate

- All plugin contributions (**views**, **pages**, **routes**, optional **models**, **migrations**, **configs**, **generators**, **DB hooks**, **CLI commands**) are declared on **`lago.Plugin`** and merged deterministically before the server or CLI runs.
- **`RegistryPlugin`** is populated when **`LoadConfigFromFile`** merges **`plugins`** (same slice you later pass to **`Start`**); there is no separate mutable **`RegistryPlugin.Register`** phase for migrated plugins.

## Application wiring

**`lago.Start`** expects the full ordered plugin list (including **`core`** if your deployment uses **`CorePlugin`**):

```go
import (
    "github.com/UniquityVentures/lago/lago"
    "github.com/UniquityVentures/lago/registry"
)

// Example: assemble once in main().
plugins := []registry.Pair[string, lago.Plugin]{
    lago.CorePlugin(db, config),       // when applicable
    p_dashboard.GetPlugin(),
    p_users.GetPlugin(),
    p_filesystem.GetPlugin(),
    // ... deployment-specific plugins returning registry.Pair[string, lago.Plugin]
}

config, err := lago.LoadConfigFromFile("example.toml", plugins)
if err != nil {
    ...
}
if err := lago.Start(config, plugins); err != nil {
    ...
}
```

**`BuildAllRegistries(plugins)`** is invoked from **`LoadConfigFromFile(path, plugins)`** (not from **`Start`**). Loading the TOML file requires **`RegistryConfig`** to already reflect the merged **`Configs`** contributions, so **`BuildAllRegistries`** runs at the beginning of **`LoadConfigFromFile`** before **`[Plugins.*]`** sections are decoded into plugin config structs and **`PostConfig`** runs.

Do **not** call **`BuildAllRegistries`** manually in **`main`** when you already use **`LoadConfigFromFile`** with the same plugin slice—it would only duplicate work. **`Start`** still takes the **`plugins`** argument so callers keep **`Load`** and **`Start`** aligned; **`Start`** does not rebuild registries.

Utility tools that only need **`LagoConfig`** fields (database connection, addresses) and never touch plugin **`[Plugins]`** tables may pass **`nil`** for **`plugins`**.

## Recommended file layout per plugin

| Concern | Where to put it |
|--------|------------------|
| **`GetPlugin`**, **`AppUrl`/`RoleUrl`** (if shared) | **`app.go`** (or **`apps.go`** if you keep the old name temporarily) |
| **`pluginViews()`** | **`views.go`** |
| **`pluginPages()`** + helpers (**`pageEntries…`**) | **`pages.go`** + **`pages_*.go`** |
| **`pluginRoutes()`** | **`paths.go`** |
| **`pluginMigrations()`** + **`migrations/`** embed | **`migrations.go`** + **`migrations/*.sql`** |
| **`pluginModels()`**, **`pluginDBInitHooks()`**, **`RegistryAdmin`** (if still **`init`**) | **`models.go`** |
| **`pluginConfigs()`** | **`config.go`** or next to auth/config types |
| **`pluginGenerators()`** | **`generator.go`** |
| **`pluginCommandFactories()`** | **`commands.go`** |

Reference implementations under **`plugins/`**:

- **`plugins/p_dashboard/`** — views, pages, routes.
- **`plugins/p_export/`** — views, pages, routes (no migrations in-tree).
- **`plugins/p_filesystem/`** — full surface: migrations, views, pages, routes, configs, generators, **`RegistryAdmin`** in **`init()`**.
- **`plugins/p_users/`** — migrations, models, configs (auth), generators, CLI, views, pages, routes, admin registration.

---

## `GetPlugin` shape

Expose a function that returns a **named pair** so the loader can key the UI and registry metadata:

```go
func GetPlugin() registry.Pair[string, lago.Plugin] {
    u, err := url.Parse(AppUrl)
    if err != nil {
        log.Panic(err)
    }

    return registry.Pair[string, lago.Plugin]{
        Key: "p_my_plugin", // stable key used in RegistryPlugin listing
        Value: lago.Plugin{
            Type:        lago.PluginTypeApp, // or Addon / Service
            Icon:        "...",
            URL:         u,
            VerboseName: "...",
            Roles:       []string{}, // optional restriction for app listing

            Migrations:  pluginMigrations(),      // optional
            Views:       pluginViews(),
            Pages:       pluginPages(),
            Routes:      pluginRoutes(),
            Models:      pluginModels(),          // optional
            Layers:      lago.PluginFeatures[views.GlobalLayer]{}, // optional
            Generators:  pluginGenerators(),      // optional
            DBInitHooks: pluginDBInitHooks(),     // optional
            Configs:     pluginConfigs(),         // optional
            CommandFactories: pluginCommandFactories(), // optional
        },
    }
}
```

Only fill the **`PluginFeatures`** fields your plugin actually needs.

---

## `PluginFeatures`: `Entries` and `Patches`

**`Entries`** append **`registry.Pair[string, T]`** with stable string keys (`"plugin.ViewName"`, `"plugin.PageKey"`, etc.).

**`Patches`** are **`registry.Pair[string, func(T) T]`** used when another plugin defines a base **`Entry`** and yours transforms it (**`Build()`** applies matching patches after clone). Prefer patches for overlays (e.g. dashboard patching **`users.LoginSuccessView`** or **`base.HomeView`**) instead of duplicating full registrations when order is guaranteed.

---

## Views

**Before:** `func init() { lago.RegistryView.Register("key", view) … }`

**After:** **`pluginViews()`** returns **`lago.PluginFeatures[*views.View]`** with **`Entries`** (and **`Patches`** if you only adjust existing views).

Handlers and **`WithLayer`** chains stay identical; only the registration mechanism changes.

---

## Pages

**Before:** `lago.RegistryPage.Register(...)` scattered across **`init()`** and helpers.

**After:**

1. Helpers return **`[]registry.Pair[string, components.PageInterface]`** (e.g. **`pageEntriesMenus()`**, **`pageEntriesTables()`**).
2. **`pluginPages()`** in **`pages.go`** appends those slices into a single **`PluginFeatures[components.PageInterface]{ Entries: … }`**.

Keeping menu vs table vs forms in separate files preserves readability.

You may keep a small **`func init()`** for side effects that are **not** on **`Plugin`** (e.g. **`components.RegistryTopbar.Register`**); prefer putting topbar registrations in **`init()` inside the page/menu file closest to that UI.

---

## Routes

**Before:** `lago.RegistryRoute.Register(...)` in **`paths.go`** **`init()`**.

**After:** **`pluginRoutes()`** on **`paths.go`** returns **`PluginFeatures[lago.Route]`** with **`Path`**, **`Handler`** (often **`lago.NewDynamicView("…")`**).

If multiple plugins expose the same **`Path`**, reconcile by merge order / deployment policy; do not rely on deprecated “register-or-patch” branching without migrating that logic to **`Patches`** or a single authoritative plugin.

---

## Migrations (`UsefulFilesystem`)

Match **`plugins/p_filesystem`** and **`plugins/p_users`**:

1. Add **`migrations/`** SQL (e.g. goose-style **`-- +goose Up`** / **`Down`** if your runner expects it).
2. **`migrations.go`**:

```go
//go:embed migrations
var migrationsFS embed.FS

func pluginMigrations() lago.PluginFeatures[lago.UsefulFilesystem] {
    return lago.PluginFeatures[lago.UsefulFilesystem]{
        Entries: []registry.Pair[string, lago.UsefulFilesystem]{
            {Key: "p_my_plugin.migrations", Value: migrationsFS},
        },
    }
}
```

3. Attach **`Migrations: pluginMigrations()`** on **`Plugin`**.

Dialect and runner details are deployment-specific; keep SQL consistent with **`DBType`** in config.

---

## Models and DB bootstrap

**`Models`** (**`PluginFeatures[any]`**): register schema types for tooling (e.g. export catalog introspection). Use **`registry.Pair`** keys like **`"p_plugin.ModelName"`** and values **`SomeModel{}`** (zero values are fine **`pointerForModel`** callers).

**`DBInitHooks`**: **`func(*gorm.DB) *gorm.DB`** run from **`InitDB`** in order. Use for:

- Seeds that assume tables exist (**after migrations applied** wherever your bootstrap runs **`InitDB`**),
- Constraints or follow-up DDL if not in goose files.

Prefer **not** coupling model registration to undefined helpers; rely on **`Models` + migrations** + explicit hooks documented in your deployment.

---

## Config

**Before:** **`lago.RegistryConfig.Register("p_plugin", cfg)`** from **`init()`**.

**After:** **`Configs: pluginConfigs()`** with **`Entries`** **`{Key: "p_plugin", Value: cfg}`** where **`cfg`** implements **`lago.Config`** (**`PostConfig()`**).

Keep **`PostConfig`** idempotent side effects predictable (filesystem paths, credential loading, etc.).

---

## Generators

**Before:** **`lago.RegistryGenerator.Register`**.

**After:** **`Generators`** with **`pluginGenerators()`** and **`Entries`** keyed like **`"plugin.Generator"`**.

---

## CLI commands

**Before:** **`lago.RegistryCommand.Register("p_plugin.foo", factory)`**.

**After:** **`CommandFactories`** (**`pluginCommandFactories()`**) with **`lago.CommandFactory`** values (**`func(lago.LagoConfig) *cobra.Command`**).

---

## Admin panel (`RegistryAdmin`)

If **`RegistryAdmin`** is still a **`Registry[T]`** in your tree, **`lago.AdminPanel`** registration can remain a tiny **`func init()`** in **`models.go`** (or **`admin.go`**) separate from **`GetPlugin`**, until **`Plugin`** grows an admin **`PluginFeatures`** field.

---

## Checklist before removing old `Register`/`init`

- [ ] **`GetPlugin()`** returns **`registry.Pair[string, lago.Plugin]`** with stable **`Key`**.
- [ ] Plugin is included in **`Start(config, plugins)`** slice.
- [ ] No remaining **`RegistryView` / RegistryPage / RegistryRoute / RegistryGenerator / RegistryConfig / RegistryCommand`** **`Register`** for keys owned by this plugin (unless intentional global side-effects like topbar **`init`**).
- [ ] **`LoadConfigFromFile(path, plugins)`** is used (not a separate manual **`BuildAllRegistries`**); plugin configs decode after the internal **`BuildAllRegistries`** inside **`LoadConfigFromFile`**.
- [ ] Smoke: list routes/views/pages in dev; run **`generate`** and CLI subcommands wired via **`CommandFactories`**.

---

## Further reading in this repo

- **`lago/registry_plugins.go`** — **`Plugin`** and **`PluginFeatures`**.
- **`lago/registry.go`** — **`BuildAllRegistries`** merge order into global registries.
- **`lago/commands.go`** — **`Start`**, and **`lago/config.go`** — **`LoadConfigFromFile`** (which invokes **`BuildAllRegistries`**).

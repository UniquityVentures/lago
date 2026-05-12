package p_dashboard

import (
	"context"

	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/plugins/p_users"
	"github.com/UniquityVentures/lago/registry"
	"github.com/UniquityVentures/lago/views"
)

func pluginViews() lago.PluginFeatures[*views.View] {
	return lago.PluginFeatures[*views.View]{
		Entries: []registry.Pair[string, *views.View]{
			{Key: "dashboard.AppsView", Value: lago.GetPageView("dashboard.AppsPage").WithLayer("users.auth", p_users.AuthenticationLayer{})},
		},
		Patches: []registry.Pair[string, func(*views.View) *views.View]{
			{Key: "users.LoginSuccessView", Value: func(_ *views.View) *views.View {
				return lago.RedirectView(lago.RoutePath("dashboard.AppsPage", nil))
			}},
			// base.HomeRoute uses view base.HomeView: send logged-in users to apps grid, others to login.
			{Key: "base.HomeView", Value: func(_ *views.View) *views.View {
				return lago.GetPageView("dashboard.HomeRedirectStub").
					WithLayer("users.optional_auth", p_users.OptionalAuthLayer{}).
					WithLayer("dashboard.home_root_redirect", lago.RedirectLayer{URLGetter: func(ctx context.Context) (string, error) {
						if p_users.UserPresentInContext(ctx) {
							return lago.RoutePath("dashboard.AppsPage", nil)(ctx)
						}
						return lago.RoutePath("users.LoginRoute", nil)(ctx)
					}})
			}},
		},
	}
}

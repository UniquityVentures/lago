package p_users

import (
	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/registry"
)

func pluginRoutes() lago.PluginFeatures[lago.Route] {
	return lago.PluginFeatures[lago.Route]{
		Entries: []registry.Pair[string, lago.Route]{
			{Key: "base.HomeRoute", Value: lago.Route{
				Path:    "/",
				Handler: lago.NewDynamicView("base.HomeView"),
			}},
			{Key: "users.ListRoute", Value: lago.Route{
				Path:    AppUrl,
				Handler: lago.NewDynamicView("users.ListView"),
			}},
			{Key: "users.CreateRoute", Value: lago.Route{
				Path:    AppUrl + "create/",
				Handler: lago.NewDynamicView("users.CreateView"),
			}},
			{Key: "users.DetailRoute", Value: lago.Route{
				Path:    AppUrl + "{id}/",
				Handler: lago.NewDynamicView("users.DetailView"),
			}},
			{Key: "users.UpdateRoute", Value: lago.Route{
				Path:    AppUrl + "{id}/edit/",
				Handler: lago.NewDynamicView("users.UpdateView"),
			}},
			{Key: "users.SelfDetailRoute", Value: lago.Route{
				Path:    AppUrl + "self/",
				Handler: lago.NewDynamicView("users.SelfDetailView"),
			}},
			{Key: "users.SelfUpdateRoute", Value: lago.Route{
				Path:    AppUrl + "self/edit/",
				Handler: lago.NewDynamicView("users.SelfUpdateView"),
			}},
			{Key: "users.SelfChangePasswordRoute", Value: lago.Route{
				Path:    AppUrl + "self/change-password/",
				Handler: lago.NewDynamicView("users.SelfChangePasswordView"),
			}},
			{Key: "users.DeleteRoute", Value: lago.Route{
				Path:    AppUrl + "{id}/delete/",
				Handler: lago.NewDynamicView("users.DeleteView"),
			}},
			{Key: "users.ChangePasswordRoute", Value: lago.Route{
				Path:    AppUrl + "{id}/change-password/",
				Handler: lago.NewDynamicView("users.ChangePasswordView"),
			}},
			{Key: "users.SelectRoute", Value: lago.Route{
				Path:    AppUrl + "select/",
				Handler: lago.NewDynamicView("users.SelectView"),
			}},
			{Key: "users.RoleSelectRoute", Value: lago.Route{
				Path:    RoleUrl + "select/",
				Handler: lago.NewDynamicView("users.RoleSelectView"),
			}},
			{Key: "users.RoleListRoute", Value: lago.Route{
				Path:    RoleUrl,
				Handler: lago.NewDynamicView("users.RoleListView"),
			}},
			{Key: "users.RoleCreateRoute", Value: lago.Route{
				Path:    RoleUrl + "create/",
				Handler: lago.NewDynamicView("users.RoleCreateView"),
			}},
			{Key: "users.RoleDetailRoute", Value: lago.Route{
				Path:    RoleUrl + "{id}/",
				Handler: lago.NewDynamicView("users.RoleDetailView"),
			}},
			{Key: "users.RoleUpdateRoute", Value: lago.Route{
				Path:    RoleUrl + "{id}/edit/",
				Handler: lago.NewDynamicView("users.RoleUpdateView"),
			}},
			{Key: "users.RoleDeleteRoute", Value: lago.Route{
				Path:    RoleUrl + "{id}/delete/",
				Handler: lago.NewDynamicView("users.RoleDeleteView"),
			}},
			{Key: "users.LoginRoute", Value: lago.Route{
				Path:    AppUrl + "login/",
				Handler: lago.NewDynamicView("users.LoginView"),
			}},
			{Key: "users.SignupRoute", Value: lago.Route{
				Path:    AppUrl + "signup/",
				Handler: lago.NewDynamicView("users.SignupView"),
			}},
			{Key: "users.LoginSuccessRoute", Value: lago.Route{
				Path:    AppUrl + "success/",
				Handler: lago.NewDynamicView("users.LoginSuccessView"),
			}},
			{Key: "users.UnauthenticatedRoute", Value: lago.Route{
				Path:    AppUrl + "unauthenticated/",
				Handler: lago.NewDynamicView("users.UnauthenticatedView"),
			}},
			{Key: "users.LogoutRoute", Value: lago.Route{
				Path:    AppUrl + "logout/",
				Handler: lago.NewDynamicView("users.LogoutView"),
			}},
		},
	}
}

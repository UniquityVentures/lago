package p_otp

import (
	"github.com/UniquityVentures/lago/components"
	"github.com/UniquityVentures/lago/lago"
	"github.com/UniquityVentures/lago/plugins/p_users"
	"github.com/UniquityVentures/lago/registry"
)

func pluginPages() lago.PluginFeatures[components.PageInterface] {
	auth := pageEntriesOtpAuth()
	prefs := pageEntriesOtpPreferences()
	entries := make([]registry.Pair[string, components.PageInterface], 0, len(auth)+len(prefs))
	entries = append(entries, auth...)
	entries = append(entries, prefs...)

	return lago.PluginFeatures[components.PageInterface]{
		Entries: entries,
		Patches: []registry.Pair[string, func(components.PageInterface) components.PageInterface]{
			{Key: "users.LoginPage", Value: patchUsersLoginPageWithOtpForgotLink},
		},
	}
}

func patchUsersLoginPageWithOtpForgotLink(page components.PageInterface) components.PageInterface {
	if scaffold, ok := page.(*components.ShellAuthScaffold); ok {
		components.InsertChildAfter(scaffold,
			"users.AuthForm",
			func(*components.FormComponent[p_users.User]) *components.ButtonLink {
				return &components.ButtonLink{
					Label: "Forgot password?",
					Link:  lago.RoutePath("otp.ForgotPasswordRoute", nil),
				}
			})
		return scaffold
	}
	panic("Base page for login page was not ShellAuthScaffold")
}

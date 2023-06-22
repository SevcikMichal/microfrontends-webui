package configuration

import (
	"os"
	"strings"
)

const (
	AcceptsLanguages           = "ACCEPTS_LANGUAGES"
	AppIconLarge               = "APP_ICON_LARGE"
	AppIconSmall               = "APP_ICON_SMALL"
	ApplicationDescription     = "APPLICATION_DESCRIPTION"
	ApplicationShellContext    = "APPLICATION_SHELL_CONTEXT"
	ApplicationTitleShort      = "APPLICATION_TITLE_SHORT"
	ApplicationTitle           = "APPLICATION_TITLE"
	BaseURL                    = "BASE_URL"
	FaviconIco                 = "FAVICON_ICO"
	ForcedRefreshPeriodSeconds = "FORCED_REFRESH_PERIOD_SECONDS"
	HttpCspHeader              = "HTTP_CSP_HEADER"
	HttpPort                   = "HTTP_PORT"
	ManifestTemplate           = "MANIFEST_TEMPLATE"
	ObserveNamespaces          = "OBSERVE_NAMESPACES"
	PwaMode                    = "PWA_MODE"
	ServiceWorker              = "SERVICE_WORKER"
	SwVersion                  = "SW_VERSION"
	SwSkipFetch                = "SW_SKIP_FETCH"
	TouchIcon                  = "TOUCH_ICON"
	UserIdHeader               = "USER_ID_HEADER"
	UserEmailHeader            = "USER_EMAIL_HEADER"
	UserNameHeader             = "USER_NAME_HEADER"
	UserRolesHeader            = "USER_ROLES_HEADER"
	WebcomponentsSelector      = "WEBCOMPONENTS_SELECTOR"
	ManifestBackgroundColor    = "MANIFEST_BACKGROUND_COLOR"
)

func GetAcceptsLanguages() []string {
	value, ok := os.LookupEnv(AcceptsLanguages)

	if ok {
		return strings.Split(value, ",")
	}

	return []string{"en"}
}

func GetAppIconLarge() string {
	value, ok := os.LookupEnv(AppIconLarge)

	if ok {
		return value
	}

	return "./assets/icon/icon.png"
}

func GetAppIconSmall() string {
	value, ok := os.LookupEnv(AppIconSmall)

	if ok {
		return value
	}

	return "./assets/icon/icon.png"
}

func GetApplicationDescription(language string) string {
	value, ok := os.LookupEnv(ApplicationDescription + "_" + strings.ToUpper(language))

	if ok {
		return value
	}

	return os.Getenv(ApplicationDescription)
}

func GetApplicationShellContext() string {
	value, ok := os.LookupEnv(ApplicationShellContext)

	if ok {
		return value
	}

	return "application-shell"
}

func GetApplicationTitleShort(language string) string {
	value, ok := os.LookupEnv(ApplicationTitleShort + "_" + strings.ToUpper(language))

	if ok {
		return value
	}

	value, ok = os.LookupEnv(ApplicationTitleShort)

	if ok {
		return value
	}

	return "Shell"
}

func GetApplicationTitle(language string) string {
	value, ok := os.LookupEnv(ApplicationTitle + "_" + strings.ToUpper(language))

	if ok {
		return value
	}

	value, ok = os.LookupEnv(ApplicationTitle)

	if ok {
		return value
	}

	return "Application shell"
}

func GetBaseURL() string {
	value, ok := os.LookupEnv(BaseURL)

	if ok {
		return value
	}

	return "/"
}

func GetFaviconIco() string {
	value, ok := os.LookupEnv(FaviconIco)

	if ok {
		return value
	}

	return "./assets/icon/favicon.ico"
}

func GetForcedRefreshPeriodSeconds() string {
	return os.Getenv(ForcedRefreshPeriodSeconds)
}

func GetHttpCspHeader() string {
	value, ok := os.LookupEnv(HttpCspHeader)

	if ok {
		return value
	}

	return "default-src 'self'; font-src 'self'; script-src 'strict-dynamic' 'nonce-{NONCE_VALUE}'; worker-src 'self'; manifest-src 'self'; style-src 'self' 'strict-dynamic';"
}

func GetHttpPort() string {
	value, ok := os.LookupEnv(HttpPort)

	if ok {
		return value
	}

	return "8082"
}

func GetManifestTemplate() string {
	value, ok := os.LookupEnv(ManifestTemplate)

	if ok {
		return value
	}

	return "manifest.template.json"
}

func GetObserveNamespaces() []string {
	value, ok := os.LookupEnv(ObserveNamespaces)

	if ok {
		return strings.Split(value, ",")
	}

	return []string{}
}

func GetPwaMode() string {
	value, ok := os.LookupEnv(PwaMode)

	if ok {
		return value
	}

	return "disabled"
}

func GetServiceWorker() string {
	value, ok := os.LookupEnv(ServiceWorker)

	if ok {
		return value
	}

	return "sw.mjs"
}

func GetSwVersion() string {
	value, ok := os.LookupEnv(SwVersion)

	if ok {
		return value
	}

	return "v1"
}

func GetSwSkipFetch() []string {
	value, ok := os.LookupEnv(SwSkipFetch)

	if ok {
		return strings.Split(value, ",")
	}

	return []string{}
}

func GetTouchIcon() string {
	value, ok := os.LookupEnv(TouchIcon)

	if ok {
		return value
	}

	return "./assets/icon/icon.png"
}

func GetUserIdHeader() string {
	value, ok := os.LookupEnv(UserIdHeader)

	if ok {
		return value
	}

	return "x-forwarded-user"
}

func GetUserEmailHeader() string {
	value, ok := os.LookupEnv(UserEmailHeader)

	if ok {
		return value
	}

	return "x-forwarded-email"
}

func GetUserNameHeader() string {
	value, ok := os.LookupEnv(UserNameHeader)

	if ok {
		return value
	}

	return "x-forwarded-preferred-username"
}

func GetUserRolesHeader() string {
	value, ok := os.LookupEnv(UserRolesHeader)

	if ok {
		return value
	}

	return "x-forwarded-groups"
}

func GetWebcomponentsSelector() []string {
	value, ok := os.LookupEnv(WebcomponentsSelector)

	if ok {
		return strings.Split(value, ",")
	}

	return []string{}
}

func GetManifestBackgroundColor() string {
	value, ok := os.LookupEnv(ManifestBackgroundColor)

	if ok {
		return value
	}

	return "#16161d"
}

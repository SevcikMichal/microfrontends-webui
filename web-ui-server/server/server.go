package server

import (
	"html/template"
	"net/http"
	"os"

	"github.com/SevcikMichal/microfrontends-webui/model"
)

func ServeSinglePageApplication(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("../web-ui/www/index.html")
	if err != nil {
		panic(err)
	}

	// Convert byte slice to string
	fileContents := string(data)

	tmpl, err := template.New("name").Parse(fileContents)
	if err != nil {
		panic(err)
	}

	pageData := &model.PageData{
		Language:                   "en",
		AppTitle:                   "Micro Frontends WebUI",
		BaseURL:                    "/",
		Description:                "Micro Frontends WebUI",
		MicroFrontendShellContext:  "application-shell",
		MicroFrontendSelector:      "",
		ProgresiveWebAppMode:       "false",
		ContentSecurityPolicyNonce: "123",
		TouchIcon:                  "",
		FavIcon:                    "",
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		panic(err)
	}
}

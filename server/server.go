package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SevcikMichal/microfrontends-webui/configuration"
	"github.com/SevcikMichal/microfrontends-webui/model"
)

func ServeSinglePageApplication(w http.ResponseWriter, r *http.Request) {
	language, _ := requestMatchLanguage(r, configuration.GetAcceptsLanguages())

	htmlToFind := "index." + language + ".html"
	matches := getAllPossibleFiles(htmlToFind)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	data, err := os.ReadFile(bestFit)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("index.html does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	fileContents := string(data)

	tmpl, err := template.New("index.html").Parse(fileContents)
	if err != nil {
		log.Panic(err)
	}

	pageData := &model.TemplateData{
		Language:                   language,
		AppTitle:                   configuration.GetApplicationTitle(language),
		BaseURL:                    configuration.GetBaseURL(),
		Description:                configuration.GetApplicationDescription(language),
		MicroFrontendShellContext:  configuration.GetApplicationShellContext(),
		MicroFrontendSelector:      strings.Join(configuration.GetWebcomponentsSelector(), ","),
		ProgresiveWebAppMode:       configuration.GetPwaMode(),
		ContentSecurityPolicyNonce: "tbd",
		TouchIcon:                  configuration.GetTouchIcon(),
		FavIcon:                    configuration.GetFaviconIco(),
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}

func ServeFile(w http.ResponseWriter, r *http.Request) {
	fileName := strings.Split(r.URL.Path, "/")[len(strings.Split(r.URL.Path, "/"))-1]

	matches := getAllPossibleFiles(fileName)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	file, err := os.Open(bestFit)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("index.html does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	http.ServeContent(w, r, matches[0], time.Now(), file)
}

func ServeManifestJson(w http.ResponseWriter, r *http.Request) {
	language, _ := requestMatchLanguage(r, configuration.GetAcceptsLanguages())
	manifest := configuration.GetManifestTemplate()
	parts := strings.Split(manifest, ".")
	ext := "." + parts[len(parts)-1]
	base := strings.Join(parts[:len(parts)-1], ".")

	manifestToFind := base + "." + language + ext

	matches := getAllPossibleFiles(manifestToFind)
	bestFit := getFirstMatchingFile("web-ui/www", matches)

	data, err := os.ReadFile(bestFit)

	if err != nil {
		if os.IsNotExist(err) {
			log.Println("manifest.template.json does not exist!")
			http.NotFound(w, r)
			return
		}
		log.Panic(err)
	}

	fileContents := string(data)

	tmpl, err := template.New("manifest.json").Parse(fileContents)
	if err != nil {
		log.Panic(err)
	}

	pageData := &model.TemplateData{
		AppTitle:        configuration.GetApplicationTitle(language),
		AppTitleShort:   configuration.GetApplicationTitleShort(language),
		BaseURL:         configuration.GetBaseURL(),
		AppIconLarge:    configuration.GetAppIconLarge(),
		AppIconSmall:    configuration.GetAppIconSmall(),
		TouchIcon:       configuration.GetTouchIcon(),
		BackgroundColor: "#ffffff",
		ThemeColor:      "#ffffff",
	}

	err = tmpl.Execute(w, pageData)
	if err != nil {
		log.Panic(err)
	}
}
